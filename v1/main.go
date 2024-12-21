package main

import (
	"fmt"
	"net/http"
	"os"
	"yuth-optician-api/config"
	"yuth-optician-api/controller"
	"yuth-optician-api/helper"
	"yuth-optician-api/models"
	"yuth-optician-api/repository"
	"yuth-optician-api/router"
	"yuth-optician-api/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	// database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("customer").AutoMigrate(models.CustomerV1{})

	// Repository
	customerRepository := repository.NewCustomerRepoImpl(db)

	// Service
	customerService := service.NewCustomerServiceImpl(customerRepository, validate)

	// Controller
	customerController := controller.NewCustomerController(customerService)

	//Router
	routes := router.NewRouter(customerController)

	listenNServe(routes)
}

func listenNServe(r *gin.Engine) {
	port := os.Getenv("PORT")
	fmt.Println(port)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	
	s := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: r,
	}
	err := s.ListenAndServe() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	helper.ErrorPanic(err)
	// fmt.Printf("Listen and serve on localhost:%s", port)
	log.Logger.Info().Msg("Started Server!")
}
