package main

import (
	"fmt"
	"math/rand"
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

func newAccount(login, password, url string) *account {
	return &account{
		url:      url,
		password: password,
		login:    login,
	}
}

func main() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1 := newAccount(login, password, url)

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
