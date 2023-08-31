package service

import (
	"rest/internal/models"
	"rest/internal/repository"

	"go.uber.org/zap"
)

type UserService struct {
	repo   *repository.Repository
	logger *zap.SugaredLogger
}

func NewUserService(repo *repository.Repository, logger *zap.SugaredLogger) *UserService {
	return &UserService{repo: repo, logger: logger}
}

func (s *UserService) CreateUser(user models.User) error {
	return s.repo.User.CreateUser(user)
}

func (s *UserService) GetUser(user models.User) (models.User, error) {
	return s.repo.User.GetUser(user)
}

func (s *UserService) UpdateUser(user models.User) error {
	return s.repo.User.UpdateUser(user)
}

func (s *UserService) DeleteUser(user models.User) error {
	return s.repo.User.DeleteUser(user)
}
