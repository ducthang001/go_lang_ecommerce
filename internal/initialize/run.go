package initialize

import (
	"fmt"

	"github.com/ducthang001/go-ecommerce-backend-api/global"
)

func Run() {
	// load configguration
	LoagConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}