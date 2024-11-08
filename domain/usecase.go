package domain

import (
	request_model "app/domain/model/request"
	"context"
	"net/url"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
)

type MemberAppUsecase interface {
	Login(ctx context.Context, payload request_model.LoginRequest) response.Base
	Register(ctx context.Context, payload request_model.RegisterRequest) response.Base
	GetMe(ctx context.Context, claim JWTClaimUser) response.Base

	SampleUserList(ctx context.Context, claim JWTClaimUser, query url.Values) response.Base
	SampleUserDetail(ctx context.Context, claim JWTClaimUser, id string) response.Base
	SampleUserExport(ctx context.Context, claim JWTClaimUser, query url.Values) response.Base
}

type AuthAppUsecase interface {
	Login(ctx context.Context, payload LoginUserRequest) response.Base
	Register(ctx context.Context, payload RegisterUserRequest ) response.Base
	GetMe(ctx context.Context, claim JWTClaimUser) response.Base
}

type FAQAppUsecase interface {
	CreateFaq(ctx context.Context, payload FaqRequest) response.Base
	GetAllFaqs(ctx context.Context) response.Base
}

type BlogAppUsecase interface {
	CreateBlog(ctx context.Context, payload BlogRequest) response.Base
	GetAllBlogs(ctx context.Context) response.Base
}
