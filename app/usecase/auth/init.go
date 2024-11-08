package usecase_auth

import (
	"app/domain"
	"app/domain/model/sql"
	jwt_helper "app/helpers/jsonwebtoken"
	"context"
	"time"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
	"gorm.io/gorm"
)

type appUsecase struct {
	mongodbRepo    domain.MongoDBRepo
	sqlRepo        *gorm.DB 
	contextTimeout time.Duration
}

type RepoInjection struct {
	MongoDBRepo domain.MongoDBRepo
	SqlDBRepo   *gorm.DB
}

func NewAppUsecase(r RepoInjection, timeout time.Duration) domain.AuthAppUsecase {
	return &appUsecase{
		sqlRepo:        r.SqlDBRepo,
		contextTimeout: timeout,
	}
}


func (u *appUsecase) Register(ctx context.Context, payload domain.RegisterUserRequest) response.Base {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	errValidation := make(map[string]string)

	// validating request
	if payload.Name == "" {
		errValidation["name"] = "name field is required"
	}

	if payload.Email == "" {
		errValidation["email"] = "email field is required"
	}

	if payload.Password == "" {
		errValidation["password"] = "password field is required"
	}

	if len(errValidation) > 0 {
		return response.ErrorValidation(errValidation, "error validation")
	}

	// check if email already exist
	result := u.sqlRepo.Where("email = ?", payload.Email)

	if result.Error == nil {
		return response.Error(400, result.Error.Error())
	}

    	if result.RowsAffected > 0 {
		return response.Error(400, "email already exist")
	}	

	// create user
	user := sql.User{
		Name:    payload.Name,
		Email:  payload.Email,
		Password: payload.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// insert user to database
	err := u.sqlRepo.Create(&user).Error

	if err != nil {
		return response.Error(400, err.Error())
	}

		// generate token
	tokenString, err := jwt_helper.GenerateJWTToken(
		jwt_helper.GetJwtCredential().Member,
		domain.JWTClaimUser{
			UserID: user.ID.String(),
		},
	)

	return response.Success(map[string]interface{}{
		"user":  user,
		"token": tokenString,
	})
}

func (u *appUsecase) Login(ctx context.Context, payload domain.LoginUserRequest) response.Base {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	errValidation := make(map[string]string)

	// validating request
	if payload.Email == "" {
		errValidation["email"] = "email field is required"
	}

	if payload.Password == "" {
		errValidation["password"] = "password field is required"
	}

	if len(errValidation) > 0 {
		return response.ErrorValidation(errValidation, "error validation")
	}

	// find user by email
	var user sql.User
	err := u.sqlRepo.Where("email = ?", payload.Email).First(&user).Error

	if err != nil {
		return response.Error(400, "user not found")
	}

	

	// check password
	if user.Password != payload.Password {
		return response.Error(400, "password not match")
	}

	// generate token
	tokenString, err := jwt_helper.GenerateJWTToken(
		jwt_helper.GetJwtCredential().Member,
		domain.JWTClaimUser{
			UserID: user.ID.String(),
		},
	)

	return response.Success(map[string]interface{}{
		"user":  user,
		"token": tokenString,
	})
}

func (u *appUsecase) GetMe(ctx context.Context, claim domain.JWTClaimUser) response.Base {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var user sql.User
	userId := claim.UserID
	result := u.sqlRepo.Where("id = ?", userId).Find(&user)
	if result.RowsAffected == 0 {
		return response.Error(404, "User not found")
	}
	if result.Error != nil {
		return response.Error(400, result.Error.Error())
	}
	

	return response.Success(user)
}