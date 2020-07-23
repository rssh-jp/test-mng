package mock

import (
	"context"
	"log"

	"github.com/rssh-jp/test-mng/api/domain"
)

var (
	hash map[string]string
)

func init() {
	hash = make(map[string]string)
}

type tokenRepository struct {
}

func NewTokenRedisMockRepository() domain.TokenRepository {
	log.Println("redis mock")
	return &tokenRepository{}
}

func (r *tokenRepository) GetByID(ctx context.Context, id string) (domain.Token, error) {
	key := id
	if token, ok := hash[key]; ok {
		return domain.Token{ID: id, Token: token}, nil
	} else {
		return domain.Token{}, domain.ErrNotFound
	}
}

func (r *tokenRepository) Store(ctx context.Context, token domain.Token) error {
	hash[token.ID] = token.Token

	return nil
}
