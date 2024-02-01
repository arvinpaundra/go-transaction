package user

import (
	"clean-arch/internal/dto"
	"clean-arch/internal/factory"
	"clean-arch/internal/model"
	"clean-arch/internal/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

type service struct {
	UserRepository repository.UserRepository
}

type Service interface {
	FindAll(ctx *gin.Context) (user []*model.User, err error)
	CreateUser(ctx *gin.Context, input dto.InsertUserRequest) (err error)
	GetById(ctx *gin.Context, ID int) (response dto.GetUserResponse, err error)
	UpdateUser(ctx *gin.Context, input dto.UpdateUserRequest, ID int) (err error)
}

// A function to call factory to initialize database connection to this/these repository
func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

// Processing the data we get from query, sometimes we don't need to process
func (s *service) FindAll(ctx *gin.Context) (user []*model.User, err error) {
	user, err = s.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.UserRepository.CurlGoogle()
	if err != nil {
		return nil, err
	}
	fmt.Println(res.Body)
	return user, nil
}

func (s *service) CreateUser(ctx *gin.Context, input dto.InsertUserRequest) (err error) {
	err = s.UserRepository.Insert(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetById(ctx *gin.Context, ID int) (response dto.GetUserResponse, err error) {
	var userResponse dto.GetUserResponse
	user, err := s.UserRepository.FindById(ctx, ID)
	if err != nil {
		return userResponse, err
	}
	userResponse = dto.GetUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return userResponse, nil
}

func (s *service) UpdateUser(ctx *gin.Context, input dto.UpdateUserRequest, ID int) (err error) {
	err = s.UserRepository.Update(ctx, input, ID)
	if err != nil {
		return err
	}
	return nil
}
