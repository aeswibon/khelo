package main

import (
	"time"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	route "github.com/cp-Coder/khelo/api/route"
	"github.com/cp-Coder/khelo/bootstrap"
	docs "github.com/cp-Coder/khelo/docs"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	docs.SwaggerInfo.BasePath = "/api"

	gin := gin.Default()
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	route.Setup(env, timeout, db, gin)
	gin.Run(env.ServerAddress)
}
