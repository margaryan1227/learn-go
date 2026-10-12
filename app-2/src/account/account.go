package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("login required")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}

	acc := Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		acc.generatePassword(12)
	}

	return &acc, nil
}

func (acc *Account) Print() {
	fmt.Println(acc.CreatedAt, " ", acc.UpdatedAt, " ", acc.Url, " ", acc.Login, " ", acc.Password)
}

func (acc *Account) generatePassword(n int) {
	pass := make([]rune, n)
	for idx := range n {
		pass[idx] = symbols[rand.IntN(len(symbols))]
	}

	acc.Password = string(pass)
}
