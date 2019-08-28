package redis

import (
	"time"
	"bytes"

	"github.com/go-redis/redis"
	"github.com/zhengchunzhe/og-rek"

	"kjlive-service/conf"
	"kjlive-service/logs"
	"strconv"
)

var (
	RC *RedisClient
)

type RedisClient struct {
	client *redis.Client
	Pong string
	err error
}

func init()  {
	RC = NewRedisClient()
}

// new a redis client
func NewRedisClient() (redisClient *RedisClient) {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Settings.RedisAddr,
		Password: conf.Settings.RedisPassword, // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()

	if err != nil {
		logs.Error(err)
	}

	return &RedisClient{client:client, Pong:pong, err:err}
}

func keyGenerate(key string) (keyG string)  {
	keyG = ":1:" + key
	return
}

// set redis key value
func Set(key string, value interface{}, expiration time.Duration)  {
	key = keyGenerate(key)
	p := &bytes.Buffer{}
	enc := ogórek.NewEncoder(p)
	err := enc.Encode(value)
	if err != nil {
		logs.Infof("key: %v | error info: %v", key, err)
	}
	err = RC.client.Set(key, p.String(), expiration).Err()
	if err != nil {
		logs.Error(err)
	}
}

// get value by key
func Get(key string) (value string) {
	key = keyGenerate(key)
	val, err := RC.client.Get(key).Bytes()
	v, err := decode(val)
	if err != nil {
		logs.Infof("key: %v | error info: %v", key, err)
	}
	if v == nil {
		return
	} else {
		return v.(string)
	}
}

// 取得数字
func GetInt(key string) (value int64) {
	key = keyGenerate(key)
	val, err := RC.client.Get(key).Bytes()
	v, err := decode(val)
	if err != nil {
		logs.Infof("key: %v | error info: %v", key, err)
	}
	if v == nil {
		return
	} else {
		switch v.(type) {
		case float64:
			return int64(v.(float64))
		case int64:
			return v.(int64)
		default:
			return
		}
	}
}


// 取得浮点数
func GetFloat(key string) (value float64) {
	key = keyGenerate(key)
	val, err := RC.client.Get(key).Bytes()
	v, err := decode(val)
	if err != nil {
		logs.Infof("key: %v | error info: %v", key, err)
	}
	if v == nil {
		return
	} else {
		return v.(float64)
	}
}


// 取得集合
func GetMap(key string) (value map[string]interface{}) {
	key = keyGenerate(key)
	val, err := RC.client.Get(key).Bytes()
	v, err := decode(val)
	if err != nil {
		logs.Infof("key: %v | error info: %v", key, err)
	}
	if v == nil {
		return
	} else {
		value = make(map[string]interface{})
		mapDecoded := v.(map[interface{}]interface{})
		for keyDecoded, valueDecoded := range mapDecoded {
			value[keyDecoded.(string)] = valueDecoded
		}
		return value
	}
}

// 取得切片数据
func GetSlice(key string) (value []interface{}) {
	key = keyGenerate(key)
	val, err := RC.client.Get(key).Bytes()
	v, err := decode(val)
	if err != nil {
		logs.Infof("key: %v | error info: %v", key, err)
	}
	if v == nil {
		return
	} else {
		return v.([]interface{})
	}
}
// get value by key
func Del(keys ...string) {
	RC.client.Del(keys...)
}

func decode(val []byte) (interface{}, error) {
	if val != nil {
		buf := bytes.NewBuffer(val)
		dec := ogórek.NewDecoder(buf)
		v, err := dec.Decode()
		return v, err
	} else {
		return nil, nil
	}
}

func GetSurveyStatus(key string) (value int) {
	key = keyGenerate(key)
	val, err := RC.client.Get(key).Result()
	if err != nil {
		return 0
	}
	res, _ := strconv.Atoi(val)
	return res
}