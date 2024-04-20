package service

import (
	"context"
	"fmt"

	"github.com/jmamadeu/episode-inviter.com/internal/data"
	"github.com/jmamadeu/episode-inviter.com/internal/model"
)

type User struct {
	db *data.Database
}

func NewUser(database *data.Database) *User {
	return &User{
		database,
	}
}

func (userService *User) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT * FROM USERS WHERE email = $1`
	var user model.User
	err := userService.db.QueryRow(ctx, query, email).Scan(&user.Id, &user.Email)
	fmt.Print(err.Error())
	if err != nil {
		return nil, err
	}

	return &user, nil
}
