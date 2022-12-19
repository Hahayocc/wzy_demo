package redis

import (
	"context"
	"time"
)

func Set(key string, value string, expireTime string) (string, error) {
	db := GetRedis()
	expTime, _ := time.ParseDuration(expireTime)
	res, err := db.Set(context.Background(), key, value, expTime).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func Get(key string) (string, error) {
	db := GetRedis()
	res, err := db.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func Del(key string) (int64, error) {
	db := GetRedis()
	res, err := db.Del(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}