package controller

import (
	"net/http"
	"strconv"
	"yuth-optician-api/data/request"
	"yuth-optician-api/data/response"
	"yuth-optician-api/helper"
	"yuth-optician-api/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CustomerController struct {
	customerService service.CustomerService
}

func NewCustomerController(service service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: service,
	}
}

/*
Create Controller
@Summary		Create customer
@Description	Save customer data in Db.
@Param			customer body request.CreateCustomerRequest true "Create customer"
@Produce		application/json
@Customer		customer
@Success		200 {object} response.Response{}
@Router			/customer [post]
*/
func (controller *CustomerController) Create(ctx *gin.Context) {
	log.Info().Msg("create tags")
	createCustomerRequest := request.CreateCustomerRequest{}
	err := ctx.ShouldBindJSON(&createCustomerRequest)
	helper.ErrorPanic(err)

	controller.customerService.Create(createCustomerRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

/*
Update Controller
@Summary		Update customer
@Description	Save tags data in Db.
@Param			tagId path string true "update customer by id"
@Param			tags body request.CreateCustomerRequest true  "Update customer"
@Produce		application/json
@Customer		customer
@Success		200 {object} response.Response{}
@Router			/customer/{customerId} [patch]
*/
func (controller *CustomerController) Update(ctx *gin.Context) {
	log.Info().Msg("update customer")
	updateCustomerRequest := request.UpdateCustomerRequest{}
	err := ctx.ShouldBindJSON(&updateCustomerRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("customerId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateCustomerRequest.Id = id

	controller.customerService.Update(updateCustomerRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

/*
Delete Controller
@Summary		Delete customer
@Description	Remove customer data by id.
@Produce		application/json
@Customer		customer
@Success		200 {object} response.Response{}
@Router			/customer/{customerID} [delete]
*/
func (controller *CustomerController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete customer")
	customerId := ctx.Param("customerId")
	id, err := strconv.Atoi(customerId)
	helper.ErrorPanic(err)
	controller.customerService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

/*
FindById Controller
@Summary		Get single customer by id.
@Description	Return the customer whos customerId value matches
@Param			customerId path string true "update customer by id"
@Produce		application/json
@Customer		customer
@Success		200 {object} response.Response{}
@Router			/customer/{customerId} [get]
*/
func (controller *CustomerController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid customer")
	customerId := ctx.Param("customerId")
	id, err := strconv.Atoi(customerId)
	helper.ErrorPanic(err)

	tagResponse := controller.customerService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

/*
Findall Controller
@Summary		Get all customers
@Description	Return list of customers.
@Customer		customer
@Success		200 {object} response.Response{}
@Router			/customer [get]
*/
func (controller *CustomerController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll customer")
	customerResponse := controller.customerService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   customerResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}