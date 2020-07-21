package domain

import (
	"context"
)

type Token struct {
	Token string `json:"token"`
}

type TokenRepository interface {
	GetByID(ctx context.Context, id string) (Token, error)
	Store(ctx context.Context, token Token) error
}
