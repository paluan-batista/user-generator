package repository

import (
	"gorm.io/gorm"
	"sync"
	"user-generator/infra/database"
	"user-generator/internal/domain/model"
)

var (
	userRepository        *UserRepository
	userRepositoryRunOnce sync.Once
)

type UserRepositoryAdapter interface {
	CreateUser(user *model.User) error
	FindUser(document string) (*model.User, error)
}

type UserRepository struct {
	repository *gorm.DB
}

func GetUserRepositoryAdapter() UserRepositoryAdapter {
	userRepositoryRunOnce.Do(func() {
		userRepository = newUserRepository()
	})

	return userRepository
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		repository: database.NewMySqlConfig(),
	}
}

func (ur UserRepository) CreateUser(user *model.User) error {
	result := ur.repository.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur UserRepository) FindUser(document string) (*model.User, error) {
	var result model.User

	user := ur.repository.Where("document = ?", document).First(&result)

	if user.Error == nil {
		return &result, nil
	}
	return nil, user.Error
}
