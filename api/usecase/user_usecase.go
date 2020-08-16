package usecase

import (
	"context"
	"math/rand"

	"github.com/rssh-jp/test-mng/api/domain"
)

const (
	validString = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`
	length      = 32
)

var (
	r *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(1))
}

type userUsecase struct {
	userRepo  domain.UserRepository
	tokenRepo domain.TokenRepository
}

func NewUserUsecase(ur domain.UserRepository, tr domain.TokenRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo:  ur,
		tokenRepo: tr,
	}
}

func (u *userUsecase) Login(ctx context.Context, id, password string) (domain.Token, error) {
	user, err := u.userRepo.GetByIDPassword(ctx, id, password)
	if err != nil {
		return domain.Token{}, err
	}

	t := newToken()
	token := domain.Token{
		ID:    user.ID,
		Token: string(t[:]),
	}

	err = u.tokenRepo.Store(ctx, token)
	if err != nil {
		return domain.Token{}, err
	}

	return token, nil
}

func (u *userUsecase) Fetch(ctx context.Context, token string) ([]domain.User, error) {
	_, err := u.tokenRepo.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	users, err := u.userRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUsecase) GetOwn(ctx context.Context, token string) (domain.User, error) {
	t, err := u.tokenRepo.GetByToken(ctx, token)
	if err != nil {
		return domain.User{}, err
	}

	user, err := u.userRepo.GetByID(ctx, t.ID)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func newToken() (ret [length]byte) {
	for i := 0; i < length; i++ {
		ret[i] = validString[r.Intn(len(validString))]
	}
	return
}
