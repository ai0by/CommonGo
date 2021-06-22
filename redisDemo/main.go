package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type RedisConfig struct {
	RedisConn redis.Conn
}

func (redisOps *RedisConfig) SetOperation(key string, value interface{}, duration int) (err error) {

	defer redisOps.redisConn.Close()

	if duration == 0 {
		_, err = redisOps.redisConn.Do("SET", key, value)
		return
	} else {
		_, err = redisOps.redisConn.Do("SETEX", key, duration, value)
		return
	}
}

func (redisOps *RedisConfig) GetOperation(key string) (value []byte, err error) {

	defer redisOps.redisConn.Close()

	value, err = redis.Bytes(redisOps.redisConn.Do("GET", key))
	return
}

func (redisOps *RedisConfig) SetJson(key string, message interface{}) (err error) {
	defer redisOps.redisConn.Close()

	var data []byte
	data, err = json.Marshal(message)

	if err != nil {
		return
	}

	_, err = redisOps.redisConn.Do("SET", key, data)

	if err != nil {
		return
	}
	return

}

func (redisOps *RedisConfig) GetJson(key string) (result []byte, err error) {
	defer redisOps.redisConn.Close()

	result, err = redis.Bytes(redisOps.redisConn.Do("GET", key))
	if err != nil {
		return
	}

	return
}

func (redisOps *RedisConfig) GetString(key string) (result string, err error) {
	defer redisOps.redisConn.Close()

	result, err = redis.String(redisOps.redisConn.Do("GET", key))

	if err != nil {
		return
	}

	return
}

func (redisOps *RedisConfig) HashGet(keyName, fieldName string) (result string, err error) {
	result, err = redis.String(redisOps.RedisConn.Do("hget", keyName, fieldName))

	if err != nil {
		return
	}

	return
}


func NewRedisConn() (*RedisConfig, error) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接redis服务器超时" + err.Error())
		return nil, err
	}
	redisOps := new(RedisConfig)
	redisOps.RedisConn = conn
	return redisOps, err
}
