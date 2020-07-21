package mysql

import (
	"database/sql"

	"github.com/rssh-jp/test-mng/api/domain"
	"github.com/rssh-jp/test-mng/api/repository/mysql/mocks"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserMysqlRepository(conn *sql.DB, opts ...Option) domain.UserRepository {
	conf := new(config)

	for _, opt := range opts {
		opt(conf)
	}

	if conf.isMock {
		return mocks.NewUserMysqlMockRepository()
	} else {
		return &userRepository{
			conn: conn,
		}
	}

}

func (r *userRepository) GetByIDPassword(ctx context.Context, id, password string) (User, error) {
	return domain.User{}, nil
}
