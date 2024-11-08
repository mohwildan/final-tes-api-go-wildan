package usecase_faq

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

type FaqAppUsecase interface {
	CreateFaq(ctx context.Context, payload domain.FaqRequest) response.Base
	GetAllFaqs(ctx context.Context) response.Base
}
func NewAppUsecase(r RepoInjection, timeout time.Duration) domain.FAQAppUsecase {
	return &appUsecase{
		sqlRepo:        r.SqlDBRepo,
		contextTimeout: timeout,
	}
}

func (a *appUsecase) CreateFaq(ctx context.Context, payload domain.FaqRequest) response.Base {
	ctx , cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()


    errValidation := make(map[string]string)

	if payload.Question == "" {
		errValidation["question"] = "question field is required"
	}

	if payload.Answer == "" {
		errValidation["answer"] = "answer field is required"
	}

	if len(errValidation) > 0 {
		return response.ErrorValidation(errValidation, "error validation")
	}

	var faq sql.Faq = sql.Faq{
		Question: payload.Question,
		Answer: payload.Answer,
	}

	result := a.sqlRepo.Create(&faq)

	if result.Error != nil {
		return response.Error(400, result.Error.Error())
	}

	return response.Success(faq)
}

func (a *appUsecase) GetAllFaqs(ctx context.Context) response.Base {
	ctx , cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	var faqs []sql.Faq

	result := a.sqlRepo.Find(&faqs)

	if result.Error != nil {
		return response.Error(400, result.Error.Error())
	}

	return response.Success(faqs)
}