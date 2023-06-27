package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	docs "github.com/cp-Coder/khelo/docs"
	"github.com/cp-Coder/khelo/internal/controllers"
	"github.com/cp-Coder/khelo/internal/middleware"
	"github.com/cp-Coder/khelo/internal/models"
	"github.com/cp-Coder/khelo/internal/routes"
	"github.com/cp-Coder/khelo/pkg/platform/cache"
	db "github.com/cp-Coder/khelo/pkg/platform/database"
	"github.com/gin-contrib/gzip"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Start the default gin server
	r := gin.Default()

	// Connecting to the redis instance
	if err := cache.InitRedis(); err != nil {
		log.Fatal("error: failed to connect to the redis database")
		return
	}
	log.Println("Successfully connected to the redis database")
	defer cache.CloseRedis()

	// Connecting to the database and migrating the models
	if err := db.OpenDBConnection(); err != nil {
		log.Fatal("error: failed to connect to the database")
		return
	}
	log.Println("Successfully connected to the database")
	defer db.CloseDBConnection()

	// Migrating the models
	userModel := &models.UserModel{}
	userModel.Init()

	authModel := &models.AuthModel{}
	authModel.Init()

	// Applying the middlewares
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	auth := &controllers.AuthController{}
	locator := &routes.ServiceLocator{
		AuthController: auth,
	}

	v1 := r.Group("/api/v1")
	routes.PublicRoutes(v1)
	routes.PrivateRoutes(v1, locator)

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	// Swagger documentation
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	if os.Getenv("SSL") == "TRUE" {
		// Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		// FIXME: Temporary fix for local development
		r.Run("127.0.0.1:" + port)
	}

}
