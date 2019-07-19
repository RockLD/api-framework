package main

import (
	"api/config"
	"api/middleware"
	"api/model"
	"api/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "path")
)

func main() {
	pflag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	//加载自定义中间件
	middlewares = append(middlewares, middleware.RequestId())
	middlewares = append(middlewares, middleware.Logging())

	router.Load(g, middlewares...)

	go func() {
		if err := pingServer(); err != nil {
			log.Info("The router bas been deployed successfully")
		}
	}()

	//初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	log.Infof("Start listen the address: %s", ":8080")

	log.Info(http.ListenAndServe(":8080", g).Error())
}

func pingServer() error {
	for i := 0; i < 2; i++ {
		resp, err := http.Get("http://127.0.0.1:8080" + "/check/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("Waiting for the router , retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
