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
			"nodes": []any{
				map[string]any{
					"id":        1,
					"location":  "us",
					"available": true,
				},
				map[string]any{
					"id":        2,
					"location":  "jp",
					"available": false,
				},
				map[string]any{
					"id":        3,
					"location":  "aus",
					"available": false,
				},
				map[string]any{
					"id":        4,
					"location":  "ru",
					"available": false,
				},
			},
			"activenode": 0,
		})
	})

	r.Run(":18888")
}
