package usecase_user

import (
	"app/domain/model/sql"
	"context"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
)



func (u *appUsecase) FindAllUser(ctx context.Context) response.Base {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var users []sql.User

	err := u.sqlRepo.Find(&users).Error

	if err != nil {
		return response.Error(400, err.Error())
	}

	return response.Success(users)
}
