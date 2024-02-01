package repository

import (
	"clean-arch/internal/dto"
	"clean-arch/internal/model"
	"clean-arch/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Define the interface for method call
type UserRepository interface {
	FindAll(ctx *gin.Context) (user []*model.User, err error)
	Insert(ctx *gin.Context, input dto.InsertUserRequest) (err error)
	FindById(ctx *gin.Context, ID int) (user *model.User, err error)
	Update(ctx *gin.Context, input dto.UpdateUserRequest, ID int) (err error)
	CurlGoogle() (*http.Response, error)
}

// Define the scoped type
type user struct {
	Db *gorm.DB
}

// Here the function is to get the database connection and to allow running the query
// This function will be called inside factory package
// gorm.DB contain database connection
func NewUserRepository(db *gorm.DB) *user {
	return &user{
		db,
	}
}

// Here we define a function that has a type *user
func (r *user) FindAll(ctx *gin.Context) (user []*model.User, err error) {
	// r was containing the database connection *gorm.DB
	// This is a gorm query builder. Read the docs, I'm too lazy to write here.
	err = r.Db.WithContext(ctx).Model(&model.User{}).Find(&user).Error
	// The most common type of error is gormErrorRecordNotFound, and gormError for the database connection
	// Most of the time we handle the record error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *user) CurlGoogle() (*http.Response, error) {
	hClient := &util.Curl{}
	res, err := hClient.To("https://google.com").Get().Do()
	if err != nil {
		return nil, err
	}
	return res, err
}

func (r *user) Insert(ctx *gin.Context, input dto.InsertUserRequest) (err error) {
	userInput := model.User{
		Name:  input.Name,
		Email: input.Email,
	}
	err = r.Db.WithContext(ctx).Create(&userInput).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *user) FindById(ctx *gin.Context, ID int) (user *model.User, err error) {
	err = r.Db.WithContext(ctx).Where("id = ?", ID).Select("id,name,email").First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *user) Update(ctx *gin.Context, input dto.UpdateUserRequest, ID int) (err error) {
	userUpdate := model.User{
		Name:  input.Name,
		Email: input.Email,
	}
	err = r.Db.WithContext(ctx).Where("id = ?", ID).Updates(&userUpdate).Error
	if err != nil {
		return err
	}
	return nil
}
