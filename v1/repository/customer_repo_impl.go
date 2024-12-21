package repository

import (
	"errors"
	"fmt"
	"yuth-optician-api/data/request"
	"yuth-optician-api/helper"
	"yuth-optician-api/models"

	"gorm.io/gorm"
)

type CustomerRepoImpl struct {
	Db *gorm.DB
}

func NewCustomerRepoImpl(Db *gorm.DB) CustomerRepository {
	return &CustomerRepoImpl{Db: Db}
}

// Delete implements CustomerRepository.
func (c *CustomerRepoImpl) Delete(customerId int) {
	var customer models.CustomerV1
	result := c.Db.Where("id = ?", customerId).Delete(&customer)
	helper.ErrorPanic(result.Error)
}

// FindAll implements CustomerRepository.
func (c *CustomerRepoImpl) FindAll() []models.CustomerV1 {
	var customers []models.CustomerV1
	result := c.Db.Find(&customers)
	helper.ErrorPanic(result.Error)
	return customers
}

// FindById implements CustomerRepository.
func (c *CustomerRepoImpl) FindById(customerId int) (customer models.CustomerV1, err error) {
	var res models.CustomerV1
	result := c.Db.Find(&res, customerId)
	if result != nil {
		return res, nil
	} else {
		var err string = fmt.Sprintf("sprintf: a %d", customerId)
		return res, errors.New(err)
	}
}

// Save implements CustomerRepository.
func (c *CustomerRepoImpl) Save(customer models.CustomerV1) {
	result := c.Db.Create(&customer)
	helper.ErrorPanic(result.Error)
}

// Update implements CustomerRepository.
func (c *CustomerRepoImpl) Update(customer models.CustomerV1) {
	var updateCustomer = request.UpdateCustomerRequest{
		Id : customer.CustomerID,
		FirstName : customer.PersonalData.FirstName,
		LastName: customer.PersonalData.LastName,
		Address: customer.PersonalData.Address,
		Phone: customer.PersonalData.Phone,
		BirthDate: customer.PersonalData.BirthDate,
		Gender: customer.PersonalData.BirthDate,
		Picture: customer.PersonalData.Picture,
	}
	result := c.Db.Model(&customer).Updates(updateCustomer)
	helper.ErrorPanic(result.Error)
}