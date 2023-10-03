package handler

import (
	"L0/internal/schema"

	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderStorage interface {
	Insert(ctx context.Context, order schema.Order, id string) error
	GetOrderById(ctx context.Context, id string) (schema.Order, error)
}

type Handler struct {
	service OrderStorage
}

func InitHandler(service OrderStorage) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Insert(c *gin.Context) {
	var order schema.Order
	err := c.BindJSON(&order)
	if err != nil {
		log.Println("Invalid order")
		c.JSON(http.StatusBadRequest, "Invalid order")

		return
	}

	err = h.service.Insert(context.TODO(), order, order.OrderUid)
	if  err != nil {
		log.Println("Can't create order")
		c.JSON(http.StatusBadRequest, "Can't create order")

		return
	}

	c.JSON(http.StatusOK, map[string]string{"uid": order.OrderUid})
}

func (h *Handler) GetOrderByid(c *gin.Context) {
	uid := c.Param("uid")
	order, err := h.service.GetOrderById(context.TODO(), uid)
	if err != nil {
		log.Println("Can't get order by id")
		c.JSON(http.StatusBadRequest, "Can't get order by id")
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("./templates/index.html")
	router.Static("/js", "./assets/js")
	router.Static("/css", "./assets/css")

	orders := router.Group("/orders")
	{
		orders.GET("", func(c *gin.Context) {
			c.HTML(200, "index.html", "")
		})
		orders.POST("", h.Insert)
		orders.GET("/:uid", h.GetOrderByid)
	}

	return router
}
