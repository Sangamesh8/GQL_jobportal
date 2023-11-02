package service

import (
	"errors"

	"gql_jobportal/graph/model"
	"gql_jobportal/repository"
)

type UserService interface {
	CreateUser(userData model.NewUser) (*model.User, error)
	CreateCompany(companyDetails model.NewCompany) (*model.Company, error)
	ViewAllCompanies() ([]*model.Company, error)
	ViewCompanyByID(companyByID string) (*model.Company, error)
	CreateJob(jobDetails model.NewJob) (*model.Job, error)
	ViewAllJob() ([]*model.Job, error)
	ViewJobByID(jobId string) (*model.Job, error)
	ViewJobByCompanyID(companyID string) ([]*model.Job, error)
}
type Service struct {
	userRepo repository.UserRepo
}

func NewService(userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be nil")
	}
	return &Service{
		userRepo: userRepo,
	}, nil

}
