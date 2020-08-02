package mock

import (
	"context"
	"log"

	"github.com/rssh-jp/test-mng/api/domain"
)

var (
	hash map[string]domain.Token
)

func init() {
	hash = make(map[string]domain.Token)
}

type tokenRepository struct {
}

func NewTokenRedisMockRepository() domain.TokenRepository {
	log.Println("redis mock")
	return &tokenRepository{}
}

func (r *tokenRepository) GetByToken(ctx context.Context, token string) (domain.Token, error) {
	key := token

	if t, ok := hash[key]; ok {
		return t, nil
	} else {
		return domain.Token{}, domain.ErrNotFound
	}
}

func (r *tokenRepository) Store(ctx context.Context, token domain.Token) error {
	hash[token.Token] = token

	return nil
}
