package usecase_user

import (
	"app/domain"
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

type UserAppUsecase interface {
	FindAllUser(ctx context.Context) response.Base
}
func NewAppUsecase(r RepoInjection, timeout time.Duration) UserAppUsecase {
	return &appUsecase{
		sqlRepo:        r.SqlDBRepo,
		contextTimeout: timeout,
	}
}