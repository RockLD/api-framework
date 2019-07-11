package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"api/router"
	"net/http"
	"time"
	"errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"api/config"
	"api/model"
	"fmt"
)

var (
	cfg = pflag.StringP("config","c","","conf/config.yaml")
)

func main() {
	pflag.Parse()




	if err := config.Init(*cfg);err != nil {
		panic(err)
	}


	for {
		fmt.Println(viper.GetString("db.addr"))
		fmt.Println(viper.GetString("db.username"))
		fmt.Println(viper.GetString("db.password"))
		fmt.Println(viper.GetInt("db.port"))
		time.Sleep(2*time.Second)
	}

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(g,middlewares...)

	go func() {
		if err := pingServer();err != nil {
			log.Info("The router bas been deployed successfully")
		}
	}()


	for {
		fmt.Println(viper.GetString("db.addr"))
		fmt.Println(viper.GetString("db.username"))
		fmt.Println(viper.GetString("db.password"))
		fmt.Println(viper.GetString("db.port"))
		time.Sleep(2*time.Second)
	}


	//初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	log.Infof("Start listen the address: %s",":8080")
	log.Info(http.ListenAndServe(":8080",g).Error())
}

func pingServer() error {
	for i := 0;i < 2;i++ {
		resp,err := http.Get("http://127.0.0.1:8080" + "/check/health")

		if err != nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("Waiting for the router , retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
