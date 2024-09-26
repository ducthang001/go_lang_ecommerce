package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	//1.
// 	sugar := zap.NewExample().Sugar()
// 	sugar.Infof("Hello name:%s, age:%d", "DucThang", 25) // like fmt.Printf()

// 	// logger
// 	logger := zap.NewExample()
// 	logger.Info("Hello", zap.String("name", "ducthang"), zap.Int("age", 25))

	//2.
	// logger := zap.NewExample()
	// logger.Info("Hello")

	// // dev env
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Development")

	// // pro env
	// logger, _ = zap.NewProduction()
	// logger.Info("Production")

	//3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// custom format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// timestamp to time
	encodeConfig.TimeKey = "time"
	//from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:22"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // .zap.AddCaller
	return zapcore.NewJSONEncoder(encodeConfig)
}

// write where?
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}