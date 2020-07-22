package usecase

import (
	"context"

	"github.com/rssh-jp/test-mng/api/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u *userUsecase) Login(ctx context.Context, id, password string) (domain.User, error) {
	user, err := u.userRepo.GetByIDPassword(ctx, id, password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
