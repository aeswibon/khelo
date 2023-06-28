package main

import (
	"time"

	"github.com/gin-gonic/gin"
	route "gitlab.com/cp-Coder/khelo/api/route"
	"gitlab.com/cp-Coder/khelo/bootstrap"
)

func main() {

	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()

	route.Setup(env, timeout, db, gin)
	gin.Run(env.ServerAddress)
}
