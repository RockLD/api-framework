package router

import (
	"apiserver/check"
	"apiserver/handler/demo"
	"apiserver/middleware"
	"github.com/gin-contrib/pprof"
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

	pprof.Register(g)
	g.POST("/login", demo.Login)

	//兼容版本，使用v1
	u := g.Group("/v1/demo")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", demo.Create)
		u.DELETE("/:id", demo.Delete)
		u.PUT("/:id", demo.Update)
		u.GET("", demo.List)
		u.GET("/:username", demo.Get)
	}

	ch := g.Group("/check")
	{
		ch.GET("/health", check.HealthCheck)
	}

	return g

}
