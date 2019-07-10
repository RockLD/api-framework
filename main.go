package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"api/router"
	"net/http"
	"time"
	"errors"
)

func main() {
	g := gin.New()
	fmt.Println("hello")

	middlewares := []gin.HandlerFunc{}

	router.Load(g,middlewares...)

	go func() {
		if err := pingServer();err != nil {
			log.Println("The router bas been deployed successfully")
		}
	}()
	
	log.Printf("Start listen the address: %s",":8080")
	log.Printf(http.ListenAndServe(":8080",g).Error())
}

func pingServer() error {
	for i := 0;i < 2;i++ {
		resp,err := http.Get("http://127.0.0.1:8080" + "/check/health")

		if err != nil && resp.StatusCode == 200 {
			return nil
		}

		log.Printf("Waiting for the router , retry in 1 second")
		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the router.")
}
