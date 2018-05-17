package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/AndreiD/BlockchainBootstrap/tools"

	"sync"
)

//the official blockchain
var Blockchain []Block

//holding incoming blocks before one is picked winner
var tempBlocks []Block

var router *gin.Engine

var bcServer chan []Block

// candidateBlocks handles incoming blocks for validation
var candidateBlocks = make(chan Block)

// announcements broadcasts winning validator to all nodes
var announcements = make(chan string)

//prevent data races
var mutex = &sync.Mutex{}

// validators keeps track of open validators and balances
var validators = make(map[string]int)

func main() {

	config, err :=  tools.ReadConfig("api_config", map[string]interface{}{
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
		genesisBlock := Block{0, time.Now().UnixNano(), "genesis", "", "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()



	bcServer = make(chan []Block)

	//broadcasting
	tick := time.NewTicker(10 * time.Second)
	go func() {
		for t := range tick.C {
			fmt.Println("Tick at", t)
		}
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
