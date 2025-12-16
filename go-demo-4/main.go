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

func main() {

	generatPassword()


	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1 := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputPassword(&account1)
}


func promptData(prompt string) string {
	fmt.Print(prompt + " ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Print(acc.login, acc.password, acc.url)
}
func generatPassword() {
	var lenghtPassword int = 10 
	for i := 0; i < lenghtPassword; i++ {
		fmt.Print(randomLetter())
	}	
}

func randomLetter() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return string(letters[rand.Intn(len(letters))])
}