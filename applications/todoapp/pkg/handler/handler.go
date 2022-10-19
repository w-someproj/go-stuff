package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/w-someproj/go-stuff/applications/todoapp/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// gin.Engine - implements interface from htttp.Handler
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// auth - for user authentification
	auth := router.Group("/auth")
	auth.POST("/sign-up", h.signUp) // registration
	auth.POST("/sign-in", h.signIn)

	// api - for working with lists and tasks
	api := router.Group("/api", h.userIdentity)
	{
		// work with lists
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById) //:id - any value that we can get by field named id
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			// work with tasks
			items := lists.Group(":id/items")
			{
				// add items to list and get all by list
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}

	return router
}
