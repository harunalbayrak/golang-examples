package main

import (
	"examples/microservices/config"
	"examples/microservices/pkg/setting"
	"examples/microservices/pkg/swagger"
	"examples/microservices/pkg/util"
	"examples/microservices/routers"
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
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
	swagger.SetupSwagger()

	config.DB, err = gorm.Open(setting.AppSettings.DBSettings.DBType, config.DBUrl())
	if err != nil {
		log.Error("Error: The database could not be opened")
	}
	defer config.DB.Close()

	util.CreateTables()

	r := routers.InitRouter()
	r.Run(fmt.Sprintf(":%s", setting.AppSettings.GeneralSettings.Port))
}
