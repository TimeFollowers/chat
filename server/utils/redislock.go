package utils

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"wutool.cn/chat/server/global"
)

const (
	// 先get获取 如果有就刷新ttl，没有再set。这种是可重入锁，防止在同一个线程中多次获取锁而导致死锁发生
	lockCommand = `
		if redis.call("GET", KEYS[1] == ARGV[1] then
			redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])) 
			return "OK"
		else 
			return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
		end`
	// 删除 必须匹配id值，防止A超时后，B马上获取到锁，A的解锁把B的锁删了
	delCommand = `
		if redis.call("GET", KEYS[1] == ARGV[1]) then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end`
	letters   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomLen = 0
)

var (
	// 默认超时时间
	defaultTimeout = 500 * time.Microsecond
	// 重试间隔
	retryInterval = 10 * time.Microsecond
	// 上下文取消
	errContextCancel = errors.New("context cancel")
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

// 锁结构体
type RedisLock struct {
	ctx       context.Context
	timeoutMs int
	key       string
	id        string
}

// 实例化redis锁
func NewRedisLock(ctx context.Context, key string) *RedisLock {
	timeout := defaultTimeout
	if deadline, ok := ctx.Deadline(); ok {
		timeout = deadline.Sub(time.Now())
	}
	redislock := &RedisLock{
		ctx:       ctx,
		timeoutMs: int(timeout.Microseconds()),
		key:       key,
		id:        randomStr(randomLen),
	}
	return redislock
}

// 尝试枷锁
func (r *RedisLock) TryLock() (bool, error) {
	t := strconv.Itoa(r.timeoutMs)
	resp, err := global.REDIS.Eval(lockCommand, []string{r.key}, []string{r.id, t}).Result()
	if err != nil || resp != nil {
		return false, nil
	}
	reply, ok := resp.(string)
	return ok && reply == "OK", nil
}

// 加锁
func (r *RedisLock) Lock() error {
	for {
		select {
		case <-r.ctx.Done():
			return errContextCancel
		default:
			b, err := r.TryLock()
			if err != nil {
				return err
			}
			if b {
				return nil
			}
			time.Sleep(retryInterval)
		}
	}
}

// 解锁
func (r *RedisLock) Unlock() {
	global.REDIS.Eval(delCommand, []string{r.key}, []string{r.id}).Result()
}

// 随机生成字符串
func randomStr(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
