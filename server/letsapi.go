package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.ebupt.com/lets/app"
)

const (
	//运行模式
	Debug  = "debug"
	Deploy = "deploy"
	Test   = "test"

	VERSION = "0.0.1"
)

var GinEngine *gin.Engine

// var LRedis interface{}

type LetsApi struct {
}

func New() *LetsApi {
	var ginRunMode string
	if app.AppConfig.RunMode == Debug {
		ginRunMode = gin.DebugMode
	} else {
		ginRunMode = gin.ReleaseMode
	}
	gin.SetMode(ginRunMode)

	GinEngine = gin.New()
	app.LLog.Info("LetsFramework.GinEngine inited")
	return &LetsApi{}
}

type routerBinder func(ginEngine *gin.Engine)

// Bind Server Router
func (l *LetsApi) BindRouter(routerBinder routerBinder) {
	routerBinder(GinEngine)
}

func (l *LetsApi) Go() {
	defer l.destruct()
	app.LLog.Info(fmt.Sprintf("LetsFramework.ApiServer run in 0.0.0.0:%v", app.AppConfig.Server.Port))
	GinEngine.Run(fmt.Sprintf("0.0.0.0:%v", app.AppConfig.Server.Port))
}

func (l LetsApi) destruct() {

	app.LLog.Info(fmt.Sprintf("LetsFramework.ApiServer 0.0.0.0:%v is stoped", app.AppConfig.Server.Port))

}
