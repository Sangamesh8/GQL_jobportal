package service

import (
	"strconv"

	"gql_jobportal/graph/model"
	"gql_jobportal/models"
)

func (s *Service) CreateUser(userData model.NewUser) (*model.User, error) {

	userDetails := models.User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: userData.Password,
	}

	userDetails, err := s.userRepo.CreateUser(userDetails)
	if err != nil {
		return nil, err
	}

	uid := strconv.FormatUint(uint64(userDetails.ID), 10)

	return &model.User{
		ID:        uid,
		Name:      userDetails.Name,
		Email:     userDetails.Email,
		CreatedAt: userDetails.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userDetails.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
