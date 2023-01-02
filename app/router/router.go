package router

import (
	"Evermos_Rakamin_BE_Project/app/config"
	"Evermos_Rakamin_BE_Project/docs"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() (err error) {
	router := gin.Default()
	router.Use(cors.Default())
	version := router.Group("/api/v1")

	gin.SetMode(gin.ReleaseMode)
	docs.SwaggerInfo.Host = "localhost" + config.AppPort()
	if config.SwaggerMode() {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		docs.SwaggerInfo.BasePath = "/api/v1"
		docsUrl := fmt.Sprintf("http://localhost:%s/swagger/index.html", config.AppPort())
		fmt.Println("\nSwagger Docs Url:", aurora.Green(docsUrl).Hyperlink(
			docsUrl,
		))
	}

	AuthRouter(version)

	err = router.Run(config.AppPort())
	return
}
