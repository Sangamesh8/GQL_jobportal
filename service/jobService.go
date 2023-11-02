package service

import (
	"gql_jobportal/graph/model"
	"gql_jobportal/models"
	"strconv"
)

func (s *Service) CreateJob(JobDetails model.NewJob) (*model.Job, error) {
	jobDetails := models.Job{
		CompanyID: JobDetails.CompanyID,
		Role:      JobDetails.Role,
		Salary:    JobDetails.Salary,
	}
	cd, err := s.userRepo.CreateJob(jobDetails)
	if err != nil {
		return nil, err
	}
	id := strconv.FormatUint(uint64(cd.ID), 10)
	return &model.Job{
		ID:        id,
		CompanyID: jobDetails.CompanyID,
		Role:      jobDetails.Role,
		Salary:    jobDetails.Salary,
	}, nil

}
func (s *Service) ViewAllJob() ([]*model.Job, error) {
	jobDetails, err := s.userRepo.ViewAllJob()
	if err != nil {
		return nil, err
	}
	var allJobs []*model.Job

	for _, v := range jobDetails {
		jobData := &model.Job{
			ID:        strconv.FormatUint(uint64(v.ID), 10),
			CompanyID: v.CompanyID,
			Role:      v.Role,
			Salary:    v.Salary,
		}
		allJobs = append(allJobs, jobData)
	}
	return allJobs, nil

}
func (s *Service) ViewJobByID(id string) (*model.Job, error) {
	jobData, err := s.userRepo.ViewByJobID(id)
	if err != nil {
		return &model.Job{}, err
	}
	return &model.Job{
		ID:        strconv.FormatUint(uint64(jobData.ID), 10),
		CompanyID: jobData.CompanyID,
		Role:      jobData.Role,
		Salary:    jobData.Salary,
	}, nil

}
func (s *Service) ViewJobByCompanyID(CompanyID string) ([]*model.Job, error) {
	jobDetails, err := s.userRepo.ViewJobByCompanyID(CompanyID)
	if err != nil {
		return nil, err
	}
	var allJobByCompany []*model.Job
	for _, v := range jobDetails {
		jobData := &model.Job{
			ID:        strconv.FormatUint(uint64(v.ID), 10),
			CompanyID: v.CompanyID,
			Role:      v.Role,
			Salary:    v.Salary,
		}
		allJobByCompany = append(allJobByCompany, jobData)
	}
	return allJobByCompany, nil

}
