package main

import (
	"demo/password/account"
	"demo/password/cloud"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// files.WriteFile("Привет Мир! я ФАЙЛ !!!!", "file.txt")
	// files.ReadFile()
	// vault := account.NewVault(files.NewJsonDb("data.json"))
	vault := account.NewVault(cloud.NewCloudDb("https://cloud.json"))
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
			createAccount(vault)
		case "2":
			color.Blue("Вы выбрали найти аккаунт")
			findAccount(vault)
		case "3":
			color.Yellow("Вы выбрали удалить аккаунт")
			deleteAccount(vault)
		case "4":
			color.Red("Вы выбрали выход")
			break menu
		}

	}

}

func deleteAccount(vault *account.VaultWithDb) {
	color.Yellow("Введите URL")
	var findUserURL string
	fmt.Scan(&findUserURL)
	countDeleted, err := vault.Vault.FindByURLtoDelete(findUserURL)
	if err != nil {
		fmt.Println("Ничего не найдено")
		return
	}
	fmt.Println("Удален ", countDeleted, " аккаунт")

}

func findAccount(vault *account.VaultWithDb) {
	color.Blue("Введите URL")
	var findUserURL string
	fmt.Scan(&findUserURL)
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

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL")
		return
	}
	color.Blue("МАЙ", vault)
	vault.AddAccount(*myAccount)

	// myAccount.OutputPassword()2
	// myAccount.GeneratPassword()
	// myAccount.OutputPassword()
}
