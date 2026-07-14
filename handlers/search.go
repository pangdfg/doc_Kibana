package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"elasticsearch-basic/config"

	"github.com/gin-gonic/gin"
)

func SearchBooks(c *gin.Context) {

	q := c.Query("q")

	if q == "" {
		c.JSON(400, gin.H{
			"error": "Query parameter q is required.",
		})
		return
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": q,
			},
		},
	}

	body, err := json.Marshal(query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res, err := config.ES.Search(
		config.ES.Search.WithContext(context.Background()),
		config.ES.Search.WithIndex("books"),
		config.ES.Search.WithBody(bytes.NewReader(body)),
		config.ES.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to search.",
		})
		return
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	hits := result["hits"].(map[string]interface{})["hits"]

	c.JSON(200, hits)
}