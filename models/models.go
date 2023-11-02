package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type Company struct {
	gorm.Model
	Name     string `json:"companyname" validate:"required"`
	Location string `json:"companylocation" validate:"required"`
}

type Job struct {
	gorm.Model
	CompanyID string `json:"companyId"`
	Role      string `json:"role"`
	Salary    string `json:"salary"`
}
