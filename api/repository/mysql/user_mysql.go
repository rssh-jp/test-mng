package mysql

import (
	"context"
	"database/sql"
	"fmt"
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
func (r *userRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	query := `
        SELECT
            id,
            name,
            age
        FROM
            users
    `

	args := []interface{}{}

	users, err := r.fetch(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return users, nil
}
func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
        UPDATE
            users
        SET
            name = ?,
            age = ?
        WHERE
            id = ?
    `

	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, user.Name, user.Age, user.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		err := fmt.Errorf("Weird Behavior. Total Affectred: %d", affect)
		return err
	}

	return nil
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
func (r *userRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	query := `
        SELECT
            id,
            name,
            age
        FROM
            users
        WHERE
            id = ?
    `

	args := []interface{}{
		id,
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
