package user

import (
	"backend-golang-evermos/src/helper"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) User
	FindAll() (users []User)
	FindById(Id int) (user User)
	FindByEmail(Email string) (user User)
	Delete(user User)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// Create is function to create data user
func (r *repository) Create(user User) User {
	err := r.db.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

// FindAll is function for get all data user
func (r *repository) FindAll() (users []User) {
	err := r.db.Find(&users).Error
	helper.PanicIfError(err)

	return
}

// FindById is function for get detail data user
func (r *repository) FindById(Id int) (user User) {
	err := r.db.Where("id = ?", Id).Find(&user).Error
	helper.PanicIfError(err)

	return
}

// FindByEmail is function for get detail data user
func (r *repository) FindByEmail(Email string) (user User) {
	err := r.db.Where("email = ?", Email).Find(&user).Error
	helper.PanicIfError(err)

	return
}

// Delete is function for delete data user
func (r *repository) Delete(user User) {
	err := r.db.Delete(&user).Error
	helper.PanicIfError(err)

	return
}
