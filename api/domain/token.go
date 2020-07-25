package domain

import (
	"context"
)

type Token struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type TokenRepository interface {
	GetByToken(ctx context.Context, token string) (Token, error)
	Store(ctx context.Context, token Token) error
}
