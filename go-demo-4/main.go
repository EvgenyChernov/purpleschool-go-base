package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	// files.WriteFile("Привет Мир! я ФАЙЛ !!!!", "file.txt")
	files.ReadFile()
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	account1, err := account.NewAccountWithTimeStam(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL")
		return
	}
	account1.OutputPassword()
	account1.GeneratPassword()
	account1.OutputPassword()

}

func promptData(prompt string) string {
	fmt.Println(prompt + " ")
	var res string
	fmt.Scan(&res)
	return res
}
