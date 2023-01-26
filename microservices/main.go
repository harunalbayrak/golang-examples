package main

import (
	"examples/microservices/config"
	"examples/microservices/docs"
	"examples/microservices/models"
	"examples/microservices/pkg/setting"
	"examples/microservices/routers"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	setting.Setup()
	fmt.Println(setting.AppSettings)

	docs.SwaggerInfo.Title = setting.AppSettings.SwaggerSettings.Title
	docs.SwaggerInfo.Description = setting.AppSettings.SwaggerSettings.Description
	docs.SwaggerInfo.Version = setting.AppSettings.SwaggerSettings.Version
	docs.SwaggerInfo.Host = setting.AppSettings.SwaggerSettings.Host
	docs.SwaggerInfo.BasePath = setting.AppSettings.SwaggerSettings.BasePath
	docs.SwaggerInfo.Schemes = setting.AppSettings.SwaggerSettings.Schemes

	config.DB, err = gorm.Open(setting.AppSettings.DBSettings.DBType, config.DBUrl())
	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Todo{})

	r := routers.InitRouter()
	r.Run(fmt.Sprintf(":%s", setting.AppSettings.Port))
}
