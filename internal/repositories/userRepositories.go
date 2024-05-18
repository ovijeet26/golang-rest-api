package repositories

import (
	"errors"
	"sync"

	"github.com/ovijeet26/golang-rest-api/internal/models"
)

type IUserRepository interface {
	Create(user models.User) error
	GetAll() ([]models.User, error)
	GetById(id string) (models.User, error)
}

type UserRepository struct {
	lock sync.Mutex
	Data map[string]models.User
}

func NewUserRepository() *UserRepository {

	return &UserRepository{
		lock: sync.Mutex{},
		Data: make(map[string]models.User, 0),
	}
}

func (ur *UserRepository) Create(user models.User) error {

	ur.lock.Lock()
	defer ur.lock.Unlock()
	if _, ok := ur.Data[user.ID]; ok {
		return errors.New("entry with the same Id already exists")
	}

	ur.Data[user.ID] = user
	return nil
}

func (ur *UserRepository) GetAll() ([]models.User, error) {

	ur.lock.Lock()
	defer ur.lock.Unlock()

	Data := make([]models.User, 0)
	if ur.Data == nil || len(ur.Data) == 0 {
		return Data, errors.New("no Data found")
	}

	for _, d := range ur.Data {
		Data = append(Data, d)
	}
	return Data, nil
}

func (ur *UserRepository) GetById(id string) (models.User, error) {

	ur.lock.Lock()
	defer ur.lock.Unlock()

	Data := models.User{}

	if _, ok := ur.Data[id]; !ok {
		return Data, errors.New("no Data found for the given ID")
	}

	Data = ur.Data[id]
	return Data, nil
}
