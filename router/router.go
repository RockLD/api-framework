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

	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
	}

	ch := g.Group("/check")
	{
		ch.GET("/health", check.HealthCheck)
	}

	return g

}
