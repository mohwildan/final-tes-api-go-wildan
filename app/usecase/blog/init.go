package usecase_blog

import (
	"app/domain"
	"app/domain/model/sql"
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

type BlogAppUsecase interface {
	CreateBlog(ctx context.Context, payload domain.BlogRequest) response.Base
	GetAllBlogs(ctx context.Context) response.Base
}

func NewAppUsecase(r RepoInjection, timeout time.Duration) domain.BlogAppUsecase {
	return &appUsecase{
		sqlRepo:        r.SqlDBRepo,
		contextTimeout: timeout,
	}
}

func (a *appUsecase) CreateBlog(ctx context.Context, payload domain.BlogRequest) response.Base {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	errValidation := make(map[string]string)

	if payload.Title == "" {
		errValidation["title"] = "title field is required"
	}

	if payload.Content == "" {
		errValidation["content"] = "content field is required"
	}

	if len(errValidation) > 0 {
		return response.ErrorValidation(errValidation, "error validation")
	}

	var blog sql.Blog = sql.Blog{
		Title:   payload.Title,
		Content: payload.Content,
		UrlImage: payload.UrlImage,
	}

	result := a.sqlRepo.Create(&blog)

	if result.Error != nil {
		return response.Error(400, result.Error.Error())
	}

	return response.Success(blog)
}

func (a *appUsecase) GetAllBlogs(ctx context.Context) response.Base {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	var blogs []sql.Blog

	result := a.sqlRepo.Find(&blogs)

	if result.Error != nil {
		return response.Error(400, result.Error.Error())
	}

	return response.Success(blogs)
}
