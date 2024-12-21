package repository

import "yuth-optician-api/models"

type CustomerRepository interface {
	Save(customer models.CustomerV1)
	Update(customer models.CustomerV1)
	Delete(customerId int)
	FindById(customerId int) (customer models.CustomerV1, err error)
	FindAll() []models.CustomerV1
}