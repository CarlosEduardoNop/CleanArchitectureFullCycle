package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"order/app/internal/order/entity"
	"order/app/internal/usecase"
)

type CreateOrderInput struct {
	Item  string  `json:"item" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func StartREST() {
	r := gin.Default()

	r.GET("/order", func(c *gin.Context) {
		service := usecase.GetOrderService{}
		orders, err := service.ListOrders()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, orders)
	})

	r.POST("/order", func(c *gin.Context) {
		var input CreateOrderInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		order := entity.Order{
			Item:  input.Item,
			Price: input.Price,
		}

		service := usecase.CreateOrderService{}

		if err := service.Create(&order); err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}

		c.JSON(201, order)
	})

	err := r.Run(":8080")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Starting rest server")
}
