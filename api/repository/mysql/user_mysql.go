package mysql

import (
	"context"
	"database/sql"
	"log"

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
		return mock.NewUserMysqlMockRepository()
	} else {
		return &userRepository{
			conn: conn,
		}
	}

}

func (r *userRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.User, error) {
	rows, err := r.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	results := make([]domain.User, 0, 8)

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, user)
	}

	return results, nil
}
func (r *userRepository) GetByIDPassword(ctx context.Context, id, password string) (domain.User, error) {
	query := `
        SELECT
            id,
            name,
            age
        FROM
            users
        WHERE
            id = ?
        AND
            password = ?
    `

	args := []interface{}{
		id,
		password,
	}

	users, err := r.fetch(ctx, query, args...)
	if err != nil {
		return domain.User{}, err
	}

	if len(users) != 1 {
		return domain.User{}, domain.ErrInvalid
	}

	return users[0], nil
}
