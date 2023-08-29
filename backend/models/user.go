package models

import (
	"backend/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(user User) (int64, error) {
	db, _ := database.InitDB()
	defer db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	res, err := db.Exec("INSERT INTO users(username, password) VALUES (?, ?)", user.Username, hashedPassword)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func Authenticate(username, password string) (*User, error) {
	db, _ := database.InitDB()
	defer db.Close()

	user := &User{}
	err := db.QueryRow("SELECT id, password FROM users WHERE username=?", username).Scan(&user.ID, &user.Password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
