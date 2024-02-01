package user_test

import (
	"clean-arch/internal/app/user"
	"clean-arch/internal/dto"
	"clean-arch/internal/factory"
	"clean-arch/internal/model"
	"clean-arch/internal/repository/mocks"
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestTable struct {
	name string
	fn   func(t *testing.T)
}

var (
	userRepository mocks.UserRepository

	service user.Service

	userModel model.User

	userReq dto.InsertUserRequest

	ctx *gin.Context
)

func initService() {
	f := factory.Factory{
		UserRepository: &userRepository,
	}

	service = user.NewService(&f)

	userModel = model.User{
		ID:    1,
		Name:  "test",
		Email: "test@mail.com",
	}

	userReq = dto.InsertUserRequest{
		Name:  "test",
		Email: "test@mail.com",
	}

	ctx = &gin.Context{}
}

func TestMain(t *testing.M) {

	initService()

	t.Run()
}

func TestCreateUser(t *testing.T) {
	tests := []TestTable{
		{
			name: "success",
			fn: func(t *testing.T) {
				userRepository.On("Insert", ctx, mock.Anything).Return(nil).Once()

				err := service.CreateUser(ctx, userReq)

				assert.Nil(t, err)
			},
		},
		{
			name: "error",
			fn: func(t *testing.T) {
				userRepository.On("Insert", ctx, mock.Anything).Return(errors.New("failed")).Once()

				err := service.CreateUser(ctx, userReq)

				assert.NotNil(t, err)
				assert.EqualError(t, err, "failed")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
