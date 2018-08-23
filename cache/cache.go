package cache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"time"
)

type Client interface {
	Insert(id string, object interface{}) error
	Get(id string, interfaceType interface{}) error
	Remove(id string) error
	SetTTL(id string, seconds int) error
	Close()
}

type RedisClient struct {
	client *redis.Client
	key    string
}

func CreateLocalRedisClient(key string) *RedisClient {
	return &RedisClient{client: redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}), key: key}
}

func CreateRedisClient(host string, port string, password string, key string) *RedisClient {
	return &RedisClient{client: redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // no password set
		DB:       0,        // use default DB
	}), key: key}
}

func (redisClient RedisClient) Insert(id string, object interface{}) error {
	err := redisClient.client.HSet(buildKey(redisClient.key, id), id, object).Err()
	if err != nil {
		errors.Wrapf(err, "failed to insert %s to %s", id, redisClient.key)
	}
	return nil
}

func (redisClient RedisClient) Get(id string, interfaceType interface{}) error {
	bytes, err := redisClient.client.HGet(buildKey(redisClient.key, id), id).Bytes()
	if err != nil {
		// error occurs if record cannot be found
		return nil
	}
	err = json.Unmarshal(bytes, interfaceType)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal %s", id)
	}
	return nil
}

func (redisClient RedisClient) Remove(id string) error {
	err := redisClient.client.HDel(buildKey(redisClient.key, id), id).Err()
	if err != nil {
		return errors.Wrapf(err, "failed to remove %s from %s", id, redisClient.key)
	}
	return nil
}

func (redisClient RedisClient) Close() {
	redisClient.client.Close()
}

func (redisClient RedisClient) SetTTL(id string, seconds int) error {
	err := redisClient.client.Expire(buildKey(redisClient.key, id), time.Duration(seconds)*time.Second).Err()
	if err != nil {
		return errors.Wrapf(err, "failed to set the till on %s to %d seconds", id, seconds)
	}
	return nil
}

func buildKey(key, id string) string {
	return key + "." + id
}
