package usecase

import (
	"context"

	"github/rssh-jp/api-test/go/domain"
)

type userUsercase struct {
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
