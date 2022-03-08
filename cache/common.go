package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisDbName string
)

func init() {
	//本地读取环境
	file, err := ini.Load("./conf/config.ini") //加载配置ini
	if err != nil {
		fmt.Println("Redis ini Load err:", err)
	}
	LoadRedis(file) //读取配置ini
	Redis()         //redis连接
}

func LoadRedis(file *ini.File) {

	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()

}
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64) //string to uint64
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	RedisClient = client
}
