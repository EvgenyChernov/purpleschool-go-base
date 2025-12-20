package main

import (
	"demo/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// files.WriteFile("Привет Мир! я ФАЙЛ !!!!", "file.txt")
	// files.ReadFile()
	color.Green("1. Создать аккаунт")
	color.Blue("2. Найти аккаунт")
	color.Yellow("3. Уадалить аккаунт")
	color.Red("4. Выход")
	var inputUserComand string
menu:
	for {
		fmt.Scan(&inputUserComand)
		switch inputUserComand {
		case "1":
			color.Green("Вы выбрали создать аккаунт")
			createAccount()
		case "2":
			color.Blue("Вы выбрали найти аккаунт")
			findAccount()
		case "3":
			color.Yellow("Вы выбрали удалить аккаунт")
			deleteAccount()
		case "4":
			color.Red("Вы выбрали выход")
			break menu
		}

	}

}

func deleteAccount() {
	color.Yellow("Введите URL")
	var findUserURL string
	fmt.Scan(&findUserURL)
	vault := account.NewVault()
	countDeleted, err := vault.FindByURLtoDelete(findUserURL)
	if err != nil {
		fmt.Println("Ничего не найдено")
		return
	}
	fmt.Println("Удален ", countDeleted, " аккаунт")

}

func findAccount() {
	color.Blue("Введите URL")
	var findUserURL string
	fmt.Scan(&findUserURL)
	vault := account.NewVault()
	findAccounts, err := vault.FindToURL(findUserURL)
	if err != nil {
		fmt.Println("Ничего не найдено")
	}
	fmt.Println(findAccounts)
	color.Green("это все что нашлось")
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
	vault := account.NewVault()
	color.Blue("МАЙ", vault)
	vault.AddAccount(*myAccount)

	// myAccount.OutputPassword()
	// myAccount.GeneratPassword()
	// myAccount.OutputPassword()
}
