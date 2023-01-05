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
