package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

func (sc *serviceConfig) dbInit() {
	redisAddress := sc.RedisHost + ":" + sc.RedisPort

	redisPool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisAddress)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	sc.RedisConn = redisPool
}
