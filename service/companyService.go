package service

import (
	"gql_jobportal/graph/model"
	"gql_jobportal/models"
	"strconv"
)

const timeStamp = "2006-01-02 15:04:05"

func (s *Service) CreateCompany(companyDetails model.NewCompany) (*model.Company, error) {
	cd := models.Company{
		Name:     companyDetails.Name,
		Location: companyDetails.Location,
	}
	cd, err := s.userRepo.CreateCompany(cd)
	if err != nil {
		return nil, err
	}
	companyId := strconv.FormatUint(uint64(cd.ID), 10)

	return &model.Company{
		ID:        companyId,
		Name:      cd.Name,
		Location:  cd.Location,
		CreatedAt: cd.CreatedAt.Format(timeStamp),
		UpdatedAt: cd.UpdatedAt.Format(timeStamp),
	}, nil
}
func (s *Service) ViewAllCompanies() ([]*model.Company, error) {
	companyDetails, err := s.userRepo.ViewAllCompany()
	if err != nil {
		return nil, err
	}
	var allCompanies []*model.Company

	for _, v := range companyDetails {
		companyData := &model.Company{
			ID:       strconv.FormatUint(uint64(v.ID), 10),
			Name:     v.Name,
			Location: v.Location,
		}
		allCompanies = append(allCompanies, companyData)
	}
	return allCompanies, nil

}
func (s *Service) ViewCompanyByID(companyByID string) (*model.Company, error) {
	companyData, err := s.userRepo.ViewCompanyByID(companyByID)
	if err != nil {
		return &model.Company{}, err
	}
	return &model.Company{
		ID:       strconv.FormatUint(uint64(companyData.ID), 10),
		Name:     companyData.Name,
		Location: companyData.Location,
	}, nil
}
