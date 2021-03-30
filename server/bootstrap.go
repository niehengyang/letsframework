package server

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"go.ebupt.com/lets/pkg/database"
	"go.ebupt.com/lets/pkg/lconfig"
	"go.ebupt.com/lets/pkg/llog"
	"go.ebupt.com/lets/pkg/lredis"
)

func configInit(configFile string) {
	//初始化配置文件
	LConfig = lconfig.New(configFile)
	LConfig.Unmarshal(&AppConfig)
}

func loggerInit() {
	//初始化日志
	var isLogDebug bool
	if AppConfig.RunMode == Debug {
		isLogDebug = true
	} else {
		isLogDebug = false
	}

	RuntimePath, err := os.Getwd()
	if err != nil {
		panic("获取当前运行路径错误")
	}

	LLog = llog.New(fmt.Sprintf("%v/%v/%v", RuntimePath, AppConfig.Logger.LogPath, AppConfig.Logger.LogFile), isLogDebug)

	LLog.Info("LetsFramework.Logger inited")
}

func dbInit() {
	c := database.DatabaseConnection{
		IP:       AppConfig.Mysql.IP,
		Port:     AppConfig.Mysql.Port,
		User:     AppConfig.Mysql.User,
		Password: AppConfig.Mysql.Password,
		Database: AppConfig.Mysql.Database,
	}
	db, err := database.New(c)
	if err != nil {
		LLog.Error(fmt.Sprintf("连接数据库错误:%v", err))
		panic(fmt.Sprintf("连接数据库错误:%v", err))
	}
	LDB = db
	LLog.Info("LetsFramework.LDB inited")
}

func redisInit() {

	redisOptions := &redis.Options{
		Addr:     fmt.Sprintf("%v:%v", AppConfig.Redis.IP, AppConfig.Redis.Port),
		Password: AppConfig.Redis.Password,
		DB:       int(AppConfig.Redis.DB),
	}

	LRedis = lredis.NewRedisClient(redisOptions)

	pong, err := LRedis.Ping(LRedisCtx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis connect error : %v , please try again", err))
	}
	LLog.Info("get redis pong :", pong)
	LLog.Info("Framework.LRedis inited")

}
