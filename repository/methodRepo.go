package repository

import (
	"errors"

	"gql_jobportal/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(userDetails models.User) (models.User, error) {
	result := r.DB.Create(&userDetails)
	if result.Error != nil {
		return models.User{}, errors.New("could not create the records")
	}
	return userDetails, nil
}
func (r *Repo) CreateCompany(companyDetails models.Company) (models.Company, error) {
	result := r.DB.Create(&companyDetails)
	if result.Error != nil {
		return models.Company{}, errors.New("could not create the records")
	}
	return companyDetails, nil

}
func (r *Repo) ViewAllCompany() ([]models.Company, error) {
	var companyDetails []models.Company
	result := r.DB.Find(&companyDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find companies")
	}
	return companyDetails, nil
}
func (r *Repo) ViewCompanyByID(CompanyId string) (models.Company, error) {
	var companyDetailsByID models.Company
	result := r.DB.Where("id = ?", CompanyId).First(&companyDetailsByID)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("could not find the company")
	}
	return companyDetailsByID, nil
}

func (r *Repo) CreateJob(JobDetails models.Job) (models.Job, error) {
	result := r.DB.Create(&JobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("could not create job")
	}
	return JobDetails, nil
}

func (r *Repo) ViewAllJob() ([]models.Job, error) {
	var AllJobs []models.Job
	result := r.DB.Find(&AllJobs)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find job")
	}
	return AllJobs, nil
}

func (r *Repo) ViewByJobID(id string) (models.Job, error) {
	var JobByID models.Job
	result := r.DB.Where("id=?", id).First(&JobByID)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("couldn't fetch jobs by ID")
	}
	return JobByID, nil
}

func (r *Repo) ViewJobByCompanyID(companyId string) ([]models.Job, error) {
	var JobByCompany []models.Job
	result := r.DB.Find(&JobByCompany)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("couldn't fetch jobs by company Id")
	}
	return JobByCompany, nil
}
