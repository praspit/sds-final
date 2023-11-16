package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"sds-final/common"

	"github.com/gin-gonic/gin"
)

var client *http.Client

type Item struct {
	Price []int    `json:"price"`
	Count []int    `json:"count"`
	Item  []string `json:"item"`
}

func main() {
	r := gin.Default()
	client = &http.Client{}

	shopMap := map[string]string{
		"shopName": "shop1",
		"shopAddr": "shop1Addr",
	}

	r.GET("/", func(c *gin.Context) { c.JSON(200, "Hello, this is shop service") })

	r.GET("/shop", func(c *gin.Context) {
		path := "http://inventory-service-service:8003/items"
		var items Item
		err := GetJson(path, &items)
		if err != nil {
			c.JSON(500, "error")
		}
		shopMap["items"] = fmt.Sprintf("%v", items)
		c.JSON(200, common.MapResponse{Data: shopMap})
	})

	r.GET("/health-check", func(c *gin.Context) { c.JSON(200, "ok") })

	r.Run(":8000")
}

func GetJson(url string, target interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
