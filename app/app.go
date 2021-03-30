package app

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"go.ebupt.com/lets/pkg/lconfig"
	"go.ebupt.com/lets/pkg/llog"
	"go.ebupt.com/lets/server/appconfig"
)

const (
	//运行模式
	Debug  = "debug"
	Deploy = "deploy"
	Test   = "test"

	VERSION = "0.0.1"
)

var AppConfig appconfig.AppConfig
var RuntimePath string

var LConfig *lconfig.LConfig
var LLog *llog.LLog

var LDB *gorm.DB
var LRedis *redis.Client
var LRedisCtx context.Context

func Bootstrap(configFile string) {

	//init Config
	configInit(configFile)

	//init Logger
	loggerInit()

	//init Database
	dbInit()

	//init Redis
	LRedisCtx = context.Background()
	redisInit()

}
