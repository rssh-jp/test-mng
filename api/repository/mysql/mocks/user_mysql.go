package mock

import (
	"context"
	"log"

	"github.com/rssh-jp/test-mng/api/domain"
)

var (
	hash map[string]*domain.User
)

func init() {
	hash = make(map[string]*domain.User)
	hash["test:test"] = &domain.User{
		ID:   "test",
		Name: "test-name",
		Age:  32,
	}
	hash["hoge:fuga"] = &domain.User{
		ID:   "hoge",
		Name: "hoge",
		Age:  55,
	}
}

type userRepository struct {
}

func NewUserMysqlMockRepository() domain.UserRepository {
	log.Println("mysql mock")
	return &userRepository{}
}

func (r *userRepository) GetByIDPassword(ctx context.Context, id, password string) (domain.User, error) {
	key := id + ":" + password
	if user, ok := hash[key]; ok {
		return *user, nil
	} else {
		return domain.User{}, domain.ErrNotFound
	}
}
