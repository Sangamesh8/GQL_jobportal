package repository

import (
	"errors"

	"gql_jobportal/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}
type UserRepo interface {
	CreateUser(userDetails models.User) (models.User, error)
	CreateCompany(companyDetails models.Company) (models.Company, error)
	ViewAllCompany() ([]models.Company, error)
	ViewCompanyByID(companyId string) (models.Company, error)
	CreateJob(JobDetails models.Job) (models.Job, error)
	ViewByJobID(id string) (models.Job, error)
	ViewAllJob() ([]models.Job, error)
	ViewJobByCompanyID(companyId string) ([]models.Job, error)
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
