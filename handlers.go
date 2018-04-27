package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"github.com/gin-gonic/gin/binding"
	"errors"
)

func HandleGetBlockchain(c *gin.Context) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	c.Writer.Write(bytes)
}

func HandleWriteBlock(c *gin.Context) {
	var json Message
	if err := c.MustBindWith(&json, binding.JSON); err == nil {

		//validate your data here
		if err := validateTheData(json.TheData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newBlock, er := generateBlock(Blockchain[len(Blockchain)-1], json.TheData)
		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": er.Error()})
			return
		}
		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			newBlockchain := append(Blockchain, newBlock)
			replaceChain(newBlockchain)
			spew.Dump(Blockchain)
		}
		c.JSON(http.StatusCreated, gin.H{"status": "block " + strconv.Itoa(newBlock.Index) + " added"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}


func validateTheData(theData string) error {
	if theData == "" {
		return errors.New("invalid data")
	}
	return nil
}
