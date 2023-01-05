package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
)

type ShopProduct struct {
	Name    string  `json:"Name,omitempty"`
	Product Product `json:"Product"`
}

type Product struct {
	Name  string  `json:"Name,omitempty"`
	Price float32 `json:"Price"`
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/prices", func(c *gin.Context) {
		shopProducts := []ShopProduct{
			{
				Name: "Biedronka",
				Product: Product{
					Name:  "Purina kurczak 1kg",
					Price: rand.Float32() * float32(rand.Intn(60)),
				},
			},
			{
				Name: "Maxi ZOO",
				Product: Product{
					Name:  "Purina kurczak 1kg",
					Price: rand.Float32() * float32(rand.Intn(60)),
				},
			},
			{
				Name: "Lidl",
				Product: Product{
					Name:  "Purina kurczak 1kg",
					Price: rand.Float32() * float32(rand.Intn(60)),
				},
			},
			// whiskas
			{
				Name: "Biedronka",
				Product: Product{
					Name:  "Whiskas kurczak 1kg",
					Price: rand.Float32() * float32(rand.Intn(60)),
				},
			},
			{
				Name: "Maxi ZOO",
				Product: Product{
					Name:  "Whiskas kurczak 1kg",
					Price: rand.Float32() * float32(rand.Intn(60)),
				},
			},
			{
				Name: "Lidl",
				Product: Product{
					Name:  "Whiskas kurczak 1kg",
					Price: rand.Float32() * float32(rand.Intn(60)),
				},
			},
		}

		c.JSON(http.StatusOK, shopProducts)
	})
	err := r.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
