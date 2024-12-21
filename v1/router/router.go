package router

import (
	"net/http"
	"yuth-optician-api/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(customerController *controller.CustomerController) *gin.Engine {
	router := gin.Default()

	// add swagger
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/customer")
	tagsRouter.GET("", customerController.FindAll)
	tagsRouter.GET("/:customerId", customerController.FindById)
	tagsRouter.POST("", customerController.Create)
	tagsRouter.PATCH("/:customerId", customerController.Update)
	tagsRouter.DELETE("/:customerId", customerController.Delete)

	return router
}