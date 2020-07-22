package mock

import (
	"context"

	"github.com/rssh-jp/test-mng/api/domain"
)

var (
	hash map[string]domain.User
)

func init() {
	hash := make(map[string]*domain.User)
	hash["test:test"] = &domain.User{
		ID:   "test-id",
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
	return &userRepository{}
}

func (r *userRepository) GetByIDPassword(ctx context.Context, id, password string) (domain.User, error) {
	if user, ok := hash[id+":"+password]; ok {
		return user, nil
	} else {
		return domain.User{}, domain.ErrNotFound
	}
}
