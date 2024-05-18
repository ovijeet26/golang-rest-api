package services

import (
	"errors"

	"github.com/ovijeet26/golang-rest-api/internal/models"
	"github.com/ovijeet26/golang-rest-api/internal/repositories"
	"github.com/ovijeet26/golang-rest-api/internal/utils"
)

type IUserService interface {
	Create(user models.User) error
	GetAll() ([]models.User, error)
	GetById(id string) (models.User, error)
	Authenticate(username, password string) (string, error)
}

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetById(id string) (models.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Authenticate(username, password string) (string, error) {
	user, err := s.repo.GetById(username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateJWT(user.Username)
}
