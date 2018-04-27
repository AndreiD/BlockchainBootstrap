package main

func InitializeRoutes() {

	// Handle the index route
	router.GET("/", HandleGetBlockchain)

	// Writes a new block
	router.POST("/", HandleWriteBlock)
}


