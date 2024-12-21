package service

import (
	"yuth-optician-api/data/request"
	"yuth-optician-api/data/response"
)

type CustomerService interface {
	Create(customer request.CreateCustomerRequest)
	Update(customer request.UpdateCustomerRequest)
	Delete(customerId int)
	FindById(customerId int) response.CustomerResponse
	FindAll() []response.CustomerResponse
}