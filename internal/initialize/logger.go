package initialize

import (
	"github.com/ducthang001/go-ecommerce-backend-api/global"
	"github.com/ducthang001/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}