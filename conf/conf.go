package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

// InitConfig is the function for initializing config
func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Sprintf("load config err: %s", err.Error()))
	}

	fmt.Println(viper.GetString("server.port"))
}
