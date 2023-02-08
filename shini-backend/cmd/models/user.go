package models

import (
	"context"
	"shini/storage/postgres"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Login     string    `json:"login,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"-"`
}
type LoginInfo struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 9)
	u.Password = string(hashedPassword)
}

func (u *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
func (u *User) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	var id uint
	stmt := "insert into users (name, login, password, created_at, email)" +
		"values ($1, $2, $3, $4, $5) returning id"
	err := postgres.DB.QueryRow(ctx, stmt, u.Name, u.Login, u.Password, u.CreatedAt, u.Email).Scan(&id)
	if err != nil {
		return err
	}
	u.Id = id
	return nil
}

func (u *LoginInfo) GetUserByLogin(user *User) error {

	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := `select id, name, login, password from users where login = $1`

	if err := postgres.DB.QueryRow(ctx, q, u.Login).Scan(&user.Id, &user.Name, &user.Login, &user.Password); err != nil {
		return err
	}

	return nil
}
func (u *User) GetUserById() error {

	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := `select name, login, password from users where id = $1`

	if err := postgres.DB.QueryRow(ctx, q, u.Id).Scan(&u.Name, &u.Login, &u.Password); err != nil {
		return err
	}

	return nil
}
