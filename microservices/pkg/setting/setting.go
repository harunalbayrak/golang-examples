package setting

import (
	"github.com/spf13/viper"
)

type General struct {
	Key            string `mapstructure:"key"`
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	TelegramToken  string `mapstructure:"telegramToken"`
	TelegramChatId string `mapstructure:"telegramChatId"`
}

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

type Token struct {
	ApiSecretKey      string `mapstructure:"apiSecretKey"`
	TokenHourLifespan string `mapstructure:"tokenHourLifespan"`
}

type App struct {
	GeneralSettings General  `mapstructure:"general"`
	DBSettings      Database `mapstructure:"database"`
	SwaggerSettings Swagger  `mapstructure:"swagger"`
	LogSettings     Log      `mapstructure:"log"`
	TokenSettings   Token    `mapstructure:"token"`
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
