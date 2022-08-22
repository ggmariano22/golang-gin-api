package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"golang/api-go-routine/models"
	"golang/api-go-routine/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/products", createProducts)
	r.GET("/products", getProducts)

	r.Run("localhost:8001")
}

func getProducts(c *gin.Context) {
	params := c.Request.URL.Query()
	page, err := strconv.Atoi(params.Get("page"))
	limit, err := strconv.Atoi(params.Get("limit"))

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"message": "Products successfully recovered",
		"data":    service.GetProducts(page, limit).Data,
	})
}

func createProducts(c *gin.Context) {
	start := time.Now()

	var products models.Products
	if err := c.BindJSON(&products); err != nil {
		log.Panic(err)
	}

	service.CreateProducts(products)

	//Prints how much this process took to be executed
	fmt.Println("Script executado em", time.Now().Sub(start))

	c.JSON(201, gin.H{
		"message": "Products successfully created!",
	})
}
