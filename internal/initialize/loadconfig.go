package initialize

import (
	"fmt"

	"github.com/ducthang001/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoagConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path config
	viper.SetConfigName("local")     // name file
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	// read server configguration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Key::", viper.GetString("security.jwt.key"))

	// config struct
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}