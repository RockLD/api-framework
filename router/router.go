package router

import (
	"api/check"
	"api/handler/user"
	"api/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, m ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())

	g.Use(middleware.NoCache)

	g.Use(middleware.Options)

	g.Use(middleware.Secure)

	g.Use(m...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The router is not exist!")
	})

	g.POST("/login", user.Login)

	//兼容版本，使用v1
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	ch := g.Group("/check")
	{
		ch.GET("/health", check.HealthCheck)
	}

	return g

}
