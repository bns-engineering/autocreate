package config

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

//Config - global config
var Config *Configuration

//LoadConfig Function
func LoadConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println(err, "fail read config")
		os.Exit(1)
	}
	return configuration
}

func getFileName() string {

	strfilename := fmt.Sprintf("config.json")
	return strfilename
}

//SetConfig Function
func SetConfig(cfg *Configuration) {
	Config = cfg
}
