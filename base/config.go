package base

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	DB  database `json:"db"`
	Log logging  `json:"log"`
}

type database struct {
	User        string   `json:"user"`
	Pwd         string   `json:"password"`
	Host        string   `json:"host"`
	Database    string   `json:"database"`
	Params      []string `json:"params"`
	MaxIdleConn int      `json:"MaxIdleConn"`
	MaxOpenConn int      `json:"MaxOpenConn"`
}

type logging struct {
	Level string `json:"level"`
}

func InitConfig() {
	configFile := parseFlags()
	LoadConfig(configFile)
}

func parseFlags() string {
	cfgFlag := flag.String("config", "", "config file")
	flag.Parse()
	if *cfgFlag == "" {
		flag.PrintDefaults()
		log.Fatalln("missing argument '-config'")
	}
	return *cfgFlag
}

func LoadConfig(configFile string) {
	v := viper.New()
	v.AddConfigPath("")
	v.SetConfigFile(configFile)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
