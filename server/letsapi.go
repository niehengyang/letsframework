package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
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
var GinEngine *gin.Engine
var LRedis *redis.Client
var LRedisCtx context.Context

// var LRedis interface{}

type LetsApi struct {
}

func New(configFile string) *LetsApi {

	//init Config
	configInit(configFile)

	//init Logger
	loggerInit()

	//init Database
	dbInit()

	//init Redis
	LRedisCtx = context.Background()
	redisInit()

	var ginRunMode string
	if AppConfig.RunMode == Debug {
		ginRunMode = gin.DebugMode
	} else {
		ginRunMode = gin.ReleaseMode
	}
	gin.SetMode(ginRunMode)

	GinEngine = gin.New()
	LLog.Info("LetsFramework.GinEngine inited")
	return &LetsApi{}
}

type routerBinder func(ginEngine *gin.Engine)

// Bind Server Router
func (l *LetsApi) BindRouter(routerBinder routerBinder) {
	routerBinder(GinEngine)
}

func (l *LetsApi) Go() {
	defer l.destruct()
	LLog.Info(fmt.Sprintf("LetsFramework.ApiServer run in 0.0.0.0:%v", AppConfig.Server.Port))
	GinEngine.Run(fmt.Sprintf("0.0.0.0:%v", AppConfig.Server.Port))
}

func (l LetsApi) destruct() {

	LLog.Info(fmt.Sprintf("LetsFramework.ApiServer 0.0.0.0:%v is stoped", AppConfig.Server.Port))

}
