package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatPassword() {
	var lenghtPassword int = 10
	var localPassword string
	for i := 0; i < lenghtPassword; i++ {
		localPassword += randomLetter()
	}
	acc.password = localPassword
}

func newAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	return &account{
		url:      urlString,
		password: password,
		login:    login,
	}, nil
}

func main() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL")
		return
	}
	account1.outputPassword()
	account1.generatPassword()
	account1.outputPassword()
}

func promptData(prompt string) string {
	fmt.Println(prompt + " ")
	var res string
	fmt.Scan(&res)
	return res
}

func randomLetter() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return string(letters[rand.Intn(len(letters))])
}
