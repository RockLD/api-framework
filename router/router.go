package router

import (
	"github.com/gin-gonic/gin"
	"api/middleware"
	"net/http"
	"api/check"
)

func Load(g *gin.Engine,m ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())

	g.Use(middleware.NoCache)

	g.Use(middleware.Options)

	g.Use(middleware.Secure)

	g.Use(m...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"The router is not exist!")
	})

	ch := g.Group("/check")
	{
		ch.GET("/health",check.HealthCheck)
	}

	return g


}
