package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	// files.WriteFile("Привет Мир! я ФАЙЛ !!!!", "file.txt")
	// files.ReadFile()
	createAccount()
}

func promptData(prompt string) string {
	fmt.Println(prompt + " ")
	var res string
	fmt.Scan(&res)
	return res
}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL")
		return
	}
	file, err := myAccount.ToBytes() 
	if err != nil {
		fmt.Println("Не удалось перобразовать в json")
		return
	}
	files.WriteFile(file, "data.json")

	// myAccount.OutputPassword()
	// myAccount.GeneratPassword()
	// myAccount.OutputPassword()
}
