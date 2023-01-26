package setting

import (
	"time"

	"github.com/spf13/viper"
)

type Database struct {
	DBType   string `mapstructure:"dbType"`
	DBName   string `mapstructure:"dbName"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Swagger struct {
	Title       string   `mapstructure:"title"`
	Description string   `mapstructure:"description"`
	Version     string   `mapstructure:"version"`
	Host        string   `mapstructure:"host"`
	BasePath    string   `mapstructure:"basePath"`
	Schemes     []string `mapstructure:"schemes"`
}

type Log struct {
	LogLevelMin   string `mapstructure:"logLevelMin"`
	LogLevel      string `mapstructure:"logLevel"`
	LogErrorStack string `mapstructure:"logErrorStack"`
}

type Redis struct {
	Host        string        `mapstructure:"host"`
	Password    string        `mapstructure:"password"`
	MaxIdle     int           `mapstructure:"maxIdle"`
	MaxActive   int           `mapstructure:"maxActive"`
	IdleTimeout time.Duration `mapstructure:"idleTimeout"`
}

type App struct {
	Host            string   `mapstructure:"host"`
	Port            string   `mapstructure:"port"`
	DBSettings      Database `mapstructure:"database"`
	SwaggerSettings Swagger  `mapstructure:"swagger"`
	LogSettings     Log      `mapstructure:"log"`
	RedisSettings   Redis    `mapstructure:"redis"`
}

var AppSettings = &App{}

func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&AppSettings)
	if err != nil {
		return
	}
}
