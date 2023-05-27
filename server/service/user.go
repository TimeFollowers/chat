package service

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"wutool.cn/chat/server/define"
	"wutool.cn/chat/server/global"
	"wutool.cn/chat/server/module"
	"wutool.cn/chat/server/utils"
)

type UserReponse struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
}

func Login(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	password := ctx.PostForm("password")
	if userName == "" || password == "" {
		log.Printf("username:%s, password:%s", userName, password)
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码不能为空",
		})
		return
	}
	log.Printf("user_name:%s, password:%s\n", userName, password)
	u, err := module.GetUserByUsernamePassword(userName, utils.GetMd5(password))
	log.Printf("user:%v \n", u)
	if err != nil {
		log.Printf("username:%s, password:%s", userName, password)
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "登录失败",
		})
		return
	}
	token, err := utils.GenerateToken(u.Id, u.Email)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserDetail(ctx *gin.Context) {
	u, _ := ctx.Get("user_claims")
	uc := u.(*utils.UserClaims)

	user, err := module.GetUserById(uc.Id)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": user,
	})
}

func UserQuery(ctx *gin.Context) {
	userName := ctx.Query("user_name")
	if userName == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}

	user, err := module.GetUserByUserName(userName)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	log.Printf("user:%v", user)
	uc := ctx.MustGet("user_claims").(*utils.UserClaims)

	data := UserQueryResult{
		UserName: user.UserName,
		Sex:      user.Sex,
		Email:    user.Email,
		Avatar:   user.Avatar,
		IsFriend: false,
	}

	if module.IsFriend(user.Id, uc.Id) {
		data.IsFriend = true
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "数据加载成功",
		"data": data,
	})
}

func SendCode(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if email == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "邮箱不能为空",
		})
		return
	}
	cnt, err := module.GetUserCountByUserEmail(email)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return
	}
	if cnt > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "当前邮箱已被注册",
		})
		return
	}
	code := utils.GetCode()
	log.Printf("=================code:%s\n", code)
	err = utils.SendCode(email, code)
	if err != nil {
		log.Printf("[ERROR]:%v\n", err)
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"code": -1,
		// 	"msg":  "系统错误",
		// })
		// return
	}
	if err = global.REDIS.Set(define.RegisterPrefix+email, code, time.Second*time.Duration(define.ExpireTime)).Err(); err != nil {
		log.Printf("[ERROR]:%v\n", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "验证码发送成功",
	})
}

func Register(c *gin.Context) {
	code := c.PostForm("code")
	email := c.PostForm("email")
	userName := c.PostForm("user_name")
	password := c.PostForm("password")
	if code == "" || email == "" || userName == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	// 判断账号是否唯一
	count, err := module.GetUserCountByUserName(userName)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误",
		})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "账号已被注册",
		})
		return
	}
	// 验证码是否正确
	r, err := global.REDIS.Get(define.RegisterPrefix + email).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}
	if r != code {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "验证码不正确",
		})
		return
	}
	ub := &module.User{
		UserName: userName,
		Password: utils.GetMd5(password),
		Email:    email,
	}
	module.CreateUser(ub)

	token, err := utils.GenerateToken(ub.Id, ub.Email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserAdd(ctx *gin.Context) {
	userName := ctx.PostForm("user_name")
	if userName == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	u, err := module.GetUserByUserName(userName)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据查询异常",
		})
		return
	}
	uc := ctx.MustGet("user_claims").(*utils.UserClaims)
	if module.IsFriend(u.Id, uc.Id) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "互为好友，不可重复添加",
		})
		return
	}
	// 保存房间记录
	r := &module.Room{
		UserId: uc.Id,
	}
	module.InsertOneRoom(r)
	// 保存用户与房间的关联记录
	ur := &module.UserRoom{
		UserId:   uc.Id,
		RoomId:   r.Id,
		RoomType: 1,
	}
	module.InsertOneUserRoom(ur)
	ur = &module.UserRoom{
		UserId:   u.Id,
		RoomId:   r.Id,
		RoomType: 1,
	}
	module.InsertOneUserRoom(ur)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func UserDelete(ctx *gin.Context) {
	userIdstr := ctx.Query("user_id")
	userId, err := strconv.Atoi(userIdstr)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	uc := ctx.MustGet("user_claims").(*utils.UserClaims)
	// 获取房间Identity
	roomId := module.GetUserRoomId(userId, uc.Id)
	if roomId == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "不为好友关系，无需删除",
		})
		return
	}
	// 删除user_room关联关系
	module.DeleteUserRoom(roomId)
	// 删除room_basic
	module.DeleteRoom(roomId)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

func UserList(ctx *gin.Context) {
	// uc := ctx.MustGet("user_claims").(*utils.UserClaims)
	// userId := uc.Id
	users, err := module.GetUserAll()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "查询失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"users": users,
		},
	})
}

type UserQueryResult struct {
	UserName string `json:"user_name"`
	Sex      int    `bson:"sex"`
	Email    string `bson:"email"`
	Avatar   string `bson:"avatar"`
	IsFriend bool   `json:"is_friend"` // 是否是好友 【true-是，false-否】
}
