package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"pledgev2-backend/config"
	"pledgev2-backend/log"
	"time"
)

func InitRedis() *redis.Pool {
	log.Logger.Info("Init success")
	redisConf := config.Config.Redis

	RedisConn = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		Wait:        true,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConf.Address, redisConf.Port))
			if err != nil {
				return nil, err
			}
			// 密码验证
			_, err = c.Do("auth", redisConf.Password)
			if err != nil {
				panic("redis auth err " + err.Error())
			}

			// 选择db
			_, err = c.Do("select", redisConf.Db)
			if err != nil {
				panic("redis select db err " + err.Error())
			}
			return c, nil
		},
	}
	err := RedisConn.Get().Err()
	if err != nil {
		panic("redis init err" + err.Error())
	}
	return RedisConn
}

func RedisSet(key string, data interface{}, aliveSeconds int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if aliveSeconds > 0 {
		_, err = conn.Do("set", key, value, "EX", aliveSeconds)
	} else {
		_, err = conn.Do("set", key, value)
	}
	if err != nil {
		return err
	}
	return nil
}

// 获取
func RedisGet(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func RedisSetString(key string, data string, aliveSeconds int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()

	var err error
	if aliveSeconds > 0 {
		_, err = redis.String(conn.Do("set", key, data, "EX", aliveSeconds))
	} else {
		_, err = redis.String(conn.Do("set", key, data))
	}
	if err != nil {
		return err
	}
	return nil
}

func RedisGetString(key string) (string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()

	reply, err := redis.String(conn.Do("get", key))
	if err != nil {
		return "", err
	}
	return reply, nil
}

func RedisSetInt64(key string, data int64, time int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = redis.Int64(conn.Do("set", key, value))
	if err != nil {
		return err
	}
	if time != 0 {
		_, err = redis.Int64(conn.Do("expire", key, time))
		if err != nil {
			return err
		}
	}
	return nil
}

func RedisGetInt64(key string) (int64, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Int64(conn.Do("get", key))
	if err != nil {
		return -1, err
	}
	return reply, nil
}

func RedisDelete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	return redis.Bool(conn.Do("del", key))
}

func redisFlushDb() error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("flushdb")
	if err != nil {
		return err
	}
	return nil
}

func redisGetHashOne(key, name string) (interface{}, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.String(conn.Do("hgetall", key, name))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func RedisSetHash(key string, data map[string]string, time interface{}) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	for k, v := range data {
		err := conn.Send("hset", key, k, v)
		if err != nil {

		}
	}
	err := conn.Flush()
	if err != nil {
		return err
	}
	if time != nil {
		_, err = conn.Do("expire", key, time.(int))
		if err != nil {
			return err
		}
	}
	return nil
}

func RedisGetHash(key string) (map[string]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.StringMap(conn.Do("hgetall", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func RedisDelHash(key string) (bool, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("del", key)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RedisExistsHash(key string) bool {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	exists, err := redis.Bool(conn.Do("hexists", key))
	if err != nil {
		return false
	}
	return exists
}

func RedisExists(key string) bool {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	exists, err := redis.Bool(conn.Do("exists", key))
	if err != nil {
		return false
	}
	return exists
}

// TTL: Time To Live（生存时间）
func RedisGetTTL(key string) int64 {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Int64(conn.Do("ttl", key))
	if err != nil {
		return 0
	}
	return reply
}

func RedisSAdd(k, v string) int64 {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Int64(conn.Do("SADD", k, v))
	if err != nil {
		return -1
	}
	return reply
}

func RedisSmembers(k string) ([]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Strings(conn.Do("smembers", k))
	if err != nil {
		return []string{}, errors.New("redis: smembers: " + err.Error())
	}
	return reply, nil
}

type RedisEncryptionTask struct {
	RecordOrderFlowId int32  `json:"recordOrderFlow"`
	Encryption        string `json:"encryption"`
	EndTime           int64  `json:"endTime"`
}

// RedisListRpush 列表右侧添加数据
func RedisListRpush(listName string, encryption string) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("lpush", listName, encryption)
	return err
}

// 取出列表数据
func RedisListLRange(listName string) ([]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Strings(conn.Do("lrange", listName, 0, -1))
	return reply, err
}

func RedisListLRem(listName string, encryption string) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("lrem", listName, 1, encryption)
	return err
}

func RedisListLength(listName string) (interface{}, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	len, err := conn.Do("llen", listName)
	return len, err
}
