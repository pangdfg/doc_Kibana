package handlers

import (
	"bytes"
	"context"
	"encoding/json"

	"elasticsearch-basic/config"

	"github.com/gin-gonic/gin"
)

func InsertDocument(c *gin.Context) {

	index := c.Query("index")

	if index == "" {
		c.JSON(400, gin.H{
			"error": "Query parameter 'index' is required.",
		})
		return
	}

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON payload.",
		})
		return
	}

	body, err := json.Marshal(data)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := config.ES.Index(
		index,
		bytes.NewReader(body),
		config.ES.Index.WithContext(context.Background()),
	)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to insert data.",
			"error":   err.Error(),
		})
		return
	}
	defer res.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	c.JSON(200, gin.H{
		"success": true,
		"result":  result,
	})
}