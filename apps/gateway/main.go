package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"sds-final/common"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	client := &http.Client{}

	r.GET("/", func(c *gin.Context) { c.JSON(200, "Hello, this is gateway service") })
	r.GET("/health-check", func(c *gin.Context) { c.JSON(200, "ok") })

	r.GET("/shop", func(c *gin.Context) {
		shop := getAllShops(client)
		c.JSON(200, shop)
	})
}

func getAllShops(client *http.Client) map[string]map[string]string {
	path := fmt.Sprintf("http://%s:%s/shop", os.Getenv("SHOP_HOST"), os.Getenv("SHOP_PORT"))
	res, err := client.Get(path)
	if err != nil {
		return map[string]map[string]string{}
	}
	defer res.Body.Close()
	var shopRes common.MapOfMapResponse
	err = json.NewDecoder(res.Body).Decode(&shopRes)
	if err != nil {
		return map[string]map[string]string{}
	}
	return shopRes.Data
}
