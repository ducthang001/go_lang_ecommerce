package initialize

import (
	"fmt"
	"time"

	"github.com/ducthang001/go-ecommerce-backend-api/global"
	"github.com/ducthang001/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
    if err != nil {
        global.Logger.Error(errString, zap.Error(err))
        panic(err)
    }
}

func InitMysql() {
	m := global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
    db, err := gorm.Open(mysql.Open(s), &gorm.Config{
        SkipDefaultTransaction: false,
    })
    checkErrorPanic(err, "InitMySql inittialization error")
    global.Logger.Info("Initializing MySql Successfully")
    global.Mdb = db

    // set Pool 
    SetPool()
    migrateTables()
}

func SetPool() {
    m := global.Config.Mysql
    sqlDb, err := global.Mdb.DB()
    if err != nil {
        fmt.Printf("mysql error: %s::", err)
    }
    sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
    sqlDb.SetMaxOpenConns(m.MaxOpenConns)
    sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
    err := global.Mdb.AutoMigrate(
        &po.User{},
        &po.Role{},
    )
    if err != nil {
        fmt.Println("Migarting tables error::", err)
    }
}
