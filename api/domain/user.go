package domain

import (
	"context"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user *User) error
	GetByIDPassword(ctx context.Context, id, password string) (User, error)
	GetByID(ctx context.Context, id string) (User, error)
}

type UserUsecase interface {
	Login(ctx context.Context, id, password string) (Token, error)
	Fetch(ctx context.Context, token string) ([]User, error)
	Update(ctx context.Context, token string, user *User) error
	GetOwn(ctx context.Context, token string) (User, error)
}
