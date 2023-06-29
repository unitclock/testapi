package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/testapi", func(c *gin.Context) {
		c.JSON(200, map[string]any{
			"nodes": map[int]any{
				1: map[string]any{
					"location":  "us",
					"available": true,
				},
				2: map[string]any{
					"location":  "jp",
					"available": false,
				},
				03: map[string]any{
					"location":  "aus",
					"available": false,
				},
				4: map[string]any{
					"location":  "ru",
					"available": false,
				},
			},
			"activenode": 0,
		})
	})

	r.Run(":18888")
}
