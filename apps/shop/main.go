package main

import (
	"sds-final/common"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	shopMap := map[string]map[string]string{
		"shop1": {
			"shopName": "shop1",
			"shopAddr": "shop1Addr",
		},
		"shop2": {
			"shopName": "shop2",
			"shopAddr": "shop2Addr",
		},
		"shop3": {
			"shopName": "shop3",
			"shopAddr": "shop3Addr",
		},
	}

	r.GET("/", func(c *gin.Context) { c.JSON(200, "Hello, this is shop service") })

	r.GET("/shop", func(c *gin.Context) { c.JSON(200, common.MapOfMapResponse{Data: shopMap}) })

	r.GET("/shop/:shopName", func(c *gin.Context) {
		shopName := c.Param("shopName")
		c.JSON(200, common.MapResponse{Data: shopMap[shopName]})
	})

	r.GET("/health-check", func(c *gin.Context) { c.JSON(200, "ok") })

	r.Run(":8000")
}
