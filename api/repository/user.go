package repository

import (
	"context"
	"log"

	"github.com/chirag1807/websocket-go/api/model/request"
	"github.com/chirag1807/websocket-go/api/model/response"
	"github.com/chirag1807/websocket-go/error"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type AuthRepository interface {
	UserRegistration(user request.User) error
	UserLogin(user request.User) (response.User, error)
}

type authRepository struct {
	pgx *pgx.Conn
}

func NewAuthRepo(pgx *pgx.Conn) AuthRepository {
	return authRepository{
		pgx: pgx,
	}
}

func (a authRepository) UserRegistration(user request.User) error {
	var userID int64
	err := a.pgx.QueryRow(context.Background(), `INSERT INTO users (name, bio, email, password) VALUES ($1, $2, $3, $4) RETURNING id`, user.Name, user.Bio, user.Email, user.Password).Scan(&userID)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
        if ok && pgErr.Code == "23505" {
            return errorhandling.DuplicateEmailFound
        }
		log.Println(err)
		return errorhandling.RegistrationFailedError
	}
	return nil
}

func (a authRepository) UserLogin(user request.User) (response.User, error) {
	var dbUser response.User
	row := a.pgx.QueryRow(context.Background(), `SELECT id, name, bio, email, password FROM users WHERE email = $1`, user.Email)
	err := row.Scan(&dbUser.ID, &dbUser.Name, &dbUser.Bio, &dbUser.Email, &dbUser.Password)

	if err != nil && err.Error() == "no rows in result set" {
		return response.User{}, errorhandling.NoUserFound
	}

	return dbUser, nil
}