package initialize

import (
	"fmt"

	"github.com/ducthang001/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	// load configguration
	LoagConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}