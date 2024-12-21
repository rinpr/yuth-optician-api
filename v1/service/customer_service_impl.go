package service

import (
	"yuth-optician-api/data/request"
	"yuth-optician-api/data/response"
	"yuth-optician-api/helper"
	"yuth-optician-api/models"
	"yuth-optician-api/repository"

	"github.com/go-playground/validator/v10"
)

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	Validate           *validator.Validate
}

func NewCustomerServiceImpl(
	cr repository.CustomerRepository,
	validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: cr,
		Validate:           validate,
	}
}

// Create implements CustomerService.
func (c *CustomerServiceImpl) Create(customer request.CreateCustomerRequest) {
	err := c.Validate.Struct(customer)
	helper.ErrorPanic(err)
	result := models.CustomerV1{
		PersonalData: models.PersonalDataV1{
			FirstName: customer.FirstName,
		},
	}
	c.CustomerRepository.Save(result)
}

// Delete implements CustomerService.
func (c *CustomerServiceImpl) Delete(customerId int) {
	c.CustomerRepository.Delete(customerId)
}

// FindAll implements CustomerService.
func (c *CustomerServiceImpl) FindAll() []response.CustomerResponse {
	result := c.CustomerRepository.FindAll()

	var customers []response.CustomerResponse
	for _, v := range result {
		customer := response.CustomerResponse{
			Id: v.CustomerID,
			FirstName: v.PersonalData.FirstName,
		}
		customers = append(customers, customer)
	}

	return customers
}

// FindById implements CustomerService.
func (c *CustomerServiceImpl) FindById(customerId int) response.CustomerResponse {
	customerData, err := c.CustomerRepository.FindById(customerId)
	helper.ErrorPanic(err)

	customerResponse := response.CustomerResponse{
		Id: customerData.CustomerID,
		FirstName: customerData.PersonalData.FirstName,
	}
	return customerResponse
}

// Update implements CustomerService.
func (c *CustomerServiceImpl) Update(customer request.UpdateCustomerRequest) {
	customerData, err := c.CustomerRepository.FindById(customer.Id)
	helper.ErrorPanic(err)
	customerData.PersonalData.FirstName = customer.FirstName
	c.CustomerRepository.Update(customerData)
}