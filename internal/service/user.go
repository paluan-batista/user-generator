package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
	"user-generator/infra/log"
	"user-generator/internal/domain/dto"
	"user-generator/internal/domain/model"
	"user-generator/internal/repository"
)

var logger = log.NewLogger()

type UserServiceAdapter interface {
	CreateUser(ctx context.Context, userRequest *dto.UserDTO) error
}

type UserService struct {
	userRepository repository.UserRepositoryAdapter
}

func NewUserService() UserServiceAdapter {
	return &UserService{
		userRepository: repository.GetUserRepositoryAdapter(),
	}
}

func (us UserService) CreateUser(ctx context.Context, userRequest *dto.UserDTO) error {

	if user, _ := us.userRepository.FindUser(userRequest.Document); user != nil {
		return errorx.Decorate(nil, "user already exist")
	}

	user := parseUserRequestDtoToUser(userRequest)

	if createUserErr := us.userRepository.CreateUser(user); createUserErr != nil {
		logger.WithError(createUserErr).WithField("context_error", ctx)
		return createUserErr
	}

	logger.Trace("user created with success")
	return nil
}

func parseUserRequestDtoToUser(userRequest *dto.UserDTO) *model.User {
	return &model.User{
		ID:       uuid.New(),
		Name:     userRequest.Name,
		Active:   true,
		Document: userRequest.Document,
		Balance:  20,
	}
}
