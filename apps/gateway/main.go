package main

import (
	"encoding/json"
	"net/http"

	"sds-final/common"

	"github.com/gin-gonic/gin"
)

var client *http.Client

func main() {
	r := gin.Default()
	client = &http.Client{}

	r.GET("/", func(c *gin.Context) { c.JSON(200, "Hello, this is gateway service") })
	r.GET("/health-check", func(c *gin.Context) { c.JSON(200, "ok") })

	r.GET("/shop", func(c *gin.Context) {
		shop := getAllShops(client)
		c.JSON(200, shop)
	})

	r.GET("/payment", func(c *gin.Context) {
		payment := getPayment(client)
		c.JSON(200, payment)
	})

	r.Run(":8000")
}

func getAllShops(client *http.Client) map[string]string {
	path := "http://shop-service-service:8000/shop"
	var shopRes common.MapResponse
	err := GetJson(path, &shopRes)
	if err != nil {
		return map[string]string{}
	}
	return shopRes.Data
}

func getPayment(client *http.Client) map[string]int {
	path := "http://payment-service-service:8081/payment"
	var paymentRes map[string]int
	err := GetJson(path, &paymentRes)
	if err != nil {
		return map[string]int{}
	}
	return paymentRes
}

func GetJson(url string, target interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
