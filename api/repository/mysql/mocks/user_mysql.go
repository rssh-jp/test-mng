package mock

import (
	"context"
	"log"
	"strings"

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

func (r *userRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	users := make([]domain.User, 0, len(hash))

	for _, user := range hash {
		users = append(users, *user)
	}

	return users, nil
}

func (r *userRepository) GetByIDPassword(ctx context.Context, id, password string) (domain.User, error) {
	key := id + ":" + password
	if user, ok := hash[key]; ok {
		return *user, nil
	} else {
		return domain.User{}, domain.ErrNotFound
	}
}

func (r *userRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	for key, user := range hash {
		_id := strings.Split(key, ":")[0]
		if _id != id {
			continue
		}

		return *user, nil
	}
	return domain.User{}, domain.ErrNotFound
}
