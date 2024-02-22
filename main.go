package main

import (
	"meetnearme-web/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Update with your allowed origins
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	r.LoadHTMLGlob("./templates/**/*.html")
	r.Static("/static", "./static")
	r.GET("/", handlers.GetAppLayout)
	r.GET("/events", handlers.GetEventsPageContent)
	r.GET("/account", handlers.GetAccountPageContent)
	r.GET("/login", handlers.GetLoginPageContent)

	componentRouterGroup := r.Group("/components")
	{
		componentRouterGroup.GET("/login-form", handlers.GetLoginFormComponent)
	}
	r.SetTrustedProxies(nil)
	r.Run()
}
