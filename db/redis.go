package db

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ErrNil = errors.New("no matching record found in redis")
	ctx    = context.Background()
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           1,
		PoolSize:     2,               // 连接池大小
		MinIdleConns: 1,               // 最小空闲连接数
		DialTimeout:  5 * time.Second, // 连接超时
		ReadTimeout:  3 * time.Second, // 读取超时
		WriteTimeout: 3 * time.Second, // 写入超时
	})
}

// Close 关闭Redis连接
func CloseRedis() error {
	if RedisClient != nil {
		return RedisClient.Close()
	}
	return nil
}

// Ping 测试连接是否存活
func Ping() error {
	return RedisClient.Ping(ctx).Err()
}

// Get 获取字符串值
func Get(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", ErrNil
	}
	return val, err
}

// Set 设置字符串值
func Set(key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// Del 删除key
func Del(key string) error {
	return RedisClient.Del(ctx, key).Err()
}

// Exists 检查key是否存在
func Exists(key string) (bool, error) {
	result, err := RedisClient.Exists(ctx, key).Result()
	return result == 1, err
}

// Expire 设置key过期时间
func Expire(key string, expiration time.Duration) (bool, error) {
	return RedisClient.Expire(ctx, key, expiration).Result()
}

// TTL 获取key剩余过期时间
func TTL(key string) (time.Duration, error) {
	return RedisClient.TTL(ctx, key).Result()
}

// Incr 自增
func Incr(key string) (int64, error) {
	return RedisClient.Incr(ctx, key).Result()
}

// Decr 自减
func Decr(key string) (int64, error) {
	return RedisClient.Decr(ctx, key).Result()
}

// HSet 哈希设置
func HSet(key string, field string, value interface{}) error {
	return RedisClient.HSet(ctx, key, field, value).Err()
}

// HGet 哈希获取
func HGet(key string, field string) (string, error) {
	val, err := RedisClient.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", ErrNil
	}
	return val, err
}

// HGetAll 获取所有哈希字段
func HGetAll(key string) (map[string]string, error) {
	return RedisClient.HGetAll(ctx, key).Result()
}

// LPush 列表左推入
func LPush(key string, values ...interface{}) error {
	return RedisClient.LPush(ctx, key, values...).Err()
}

// RPop 列表右弹出
func RPop(key string) (string, error) {
	val, err := RedisClient.RPop(ctx, key).Result()
	if err == redis.Nil {
		return "", ErrNil
	}
	return val, err
}

// SAdd 集合添加
func SAdd(key string, members ...interface{}) error {
	return RedisClient.SAdd(ctx, key, members...).Err()
}

// SMembers 获取集合所有成员
func SMembers(key string) ([]string, error) {
	return RedisClient.SMembers(ctx, key).Result()
}

// ZAdd 有序集合添加
func ZAdd(key string, members ...redis.Z) error {
	return RedisClient.ZAdd(ctx, key, members...).Err()
}

// ZRange 有序集合范围查询
func ZRange(key string, start, stop int64) ([]string, error) {
	return RedisClient.ZRange(ctx, key, start, stop).Result()
}

// Publish 发布消息
func Publish(channel string, message interface{}) error {
	return RedisClient.Publish(ctx, channel, message).Err()
}

// Subscribe 订阅频道
func Subscribe(channels ...string) *redis.PubSub {
	return RedisClient.Subscribe(ctx, channels...)
}

// GetJson 获取JSON并反序列化
func GetJson(key string, v interface{}) error {
	data, err := Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), v)
}

// SetJson 序列化为JSON并存储
func SetJson(key string, v interface{}, expiration time.Duration) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return Set(key, data, expiration)
}

// Lock 分布式锁
func Lock(key string, value interface{}, expiration time.Duration) (bool, error) {
	return RedisClient.SetNX(ctx, key, value, expiration).Result()
}

// Unlock 释放分布式锁
func Unlock(key string) error {
	return Del(key)
}

// Eval 执行Lua脚本
func Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	return RedisClient.Eval(ctx, script, keys, args).Result()
}
