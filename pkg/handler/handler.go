package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/orejametov-dev/todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine  {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.singIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		lists := api.Group("lists")
		{
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.POST("/", h.createList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				 items.GET("/", h.getAllItems)
				 items.POST("/", h.createItem)
				 items.GET("/:item_id", h.getItemById)
				 items.PUT("/:item_id", h.updateItem)
				 items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}