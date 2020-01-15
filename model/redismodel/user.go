package redismodel

import (
	"github.com/jinzhu/gorm"
	"golearning/inits"
)

// JwtBlacklist JwtBlacklist
type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text"`
}

// GetRedisJWT 判断当前用户是否在线
func (j *JwtBlacklist) GetRedisJWT(userName string) (RedisJWT string, err error) {
	RedisJWT, err = inits.DEFAULTREDIS.Get(userName).Result()
	return RedisJWT, err
}

// SetRedisJWT 用户登录
func (j *JwtBlacklist) SetRedisJWT(userName string) (err error) {
	err = inits.DEFAULTREDIS.Set(userName, j.Jwt, 1000*1000*1000*60*60*24*7).Err()
	return err
}
