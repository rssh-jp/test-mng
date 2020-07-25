package redis

import (
	"context"

	"github.com/rssh-jp/test-mng/api/domain"
	"github.com/rssh-jp/test-mng/api/repository/redis/mocks"
)

type tokenRepository struct {
}

func NewTokenRedisRepository(opts ...Option) domain.TokenRepository {
	conf := new(config)

	for _, opt := range opts {
		opt(conf)
	}

	if conf.isMock {
		return mock.NewTokenRedisMockRepository()
	} else {
		return &tokenRepository{}
	}

}

func (r *tokenRepository) GetByToken(ctx context.Context, token string) (domain.Token, error) {
	return domain.Token{}, nil
}

func (r *tokenRepository) Store(ctx context.Context, token domain.Token) error {
	return nil
}
