package initialize

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wutool.cn/chat/server/global"
	"wutool.cn/chat/server/module/entity"
)

func InitMysql() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	// mysqlConfig := mysql.Config{
	// 	DSN:                       m.Dsn(),
	// 	DefaultStringSize:         191,   // string 类型字段的默认长度
	// 	SkipInitializeWithVersion: false, // 根据版本自动配置

	// }

	_db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, _ := _db.DB()
	// 设置数据库连接参数
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return _db
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		entity.User{},
		entity.Message{},
	)
	if err != nil {
		fmt.Println("迁移失败")
	}
	fmt.Println("迁移成功")
}
