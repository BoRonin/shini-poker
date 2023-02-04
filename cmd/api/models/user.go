package models

import (
    "context"
    "golang.org/x/crypto/bcrypt"
    "shini/storage/postgres"
    "time"
)

type User struct {
    Id uint `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Login string `json:"login"`
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
}

func (u *User) SetPassword(password string) {
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 9)
    u.Password = string(hashedPassword)
}

func (u *User) ComparePasswords(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
func (u *User) Create() error {
    ctx , cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
    defer cancel()
    var id uint
    stmt := "insert into users (name, login, password, created_at, email)" +
        "values ($1, $2, $3, $4, $5) returning id"
    err := postgres.DB.QueryRowContext(ctx, stmt, u.Name, u.Login, u.Password, u.CreatedAt, u.Email).Scan(&id)
    if err != nil {
        return err
    }
    u.Id = id
    return nil
}