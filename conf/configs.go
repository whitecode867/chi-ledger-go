package conf

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultRunMode    = "develop"
	productionRunMode = "production"
)

var Configs = viper.New()

func InitialiseConfigs() {
	Configs.AddConfigPath("conf")
	Configs.SetConfigFile("json")

	runMode := strings.ToLower(os.Getenv("RUN_MODE"))
	if runMode != productionRunMode || runMode != defaultRunMode {
		runMode = defaultRunMode
	}

	Configs.SetConfigName(runMode)
	if errRead := Configs.ReadInConfig(); errRead != nil {
		log.Println(errRead.Error())
		panic(errRead)
	}

	Configs.AutomaticEnv()
	Configs.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
