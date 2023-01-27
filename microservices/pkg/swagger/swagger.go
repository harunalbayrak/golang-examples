package swagger

import (
	"examples/microservices/docs"
	"examples/microservices/pkg/setting"
)

func SetupSwagger() {
	docs.SwaggerInfo.Title = setting.AppSettings.SwaggerSettings.Title
	docs.SwaggerInfo.Description = setting.AppSettings.SwaggerSettings.Description
	docs.SwaggerInfo.Version = setting.AppSettings.SwaggerSettings.Version
	docs.SwaggerInfo.Host = setting.AppSettings.SwaggerSettings.Host
	docs.SwaggerInfo.BasePath = setting.AppSettings.SwaggerSettings.BasePath
	docs.SwaggerInfo.Schemes = setting.AppSettings.SwaggerSettings.Schemes
}
