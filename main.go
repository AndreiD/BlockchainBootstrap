package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"./tools"

)

var Blockchain []Block
var router *gin.Engine

func main() {

	config, err :=  tools.ReadConfig("config", map[string]interface{}{
		"port":     1234,
		"hostname": "localhost",
		"auth": map[string]string{
			"username": "user",
			"password": "pass",
		},
	})
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	}

	//gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	InitializeRoutes()

	//Genesis
	go func() {
		genesisBlock := Block{0, time.Now().UnixNano(), "genesis", "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()

	server := &http.Server{
		Addr:           ":" + strconv.Itoa(config.GetInt("port")),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.SetKeepAlivesEnabled(false)
	server.ListenAndServe()

}
