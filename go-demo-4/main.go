package main

import (
	"demo/password/account"
	"demo/password/cloud"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	output.PrintError(1)
	output.PrintError("1")
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
		menuFunc := menu[inputUserComand]
		if menuFunc != nil {
			menuFunc(vault)
		} else {
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
	findAccounts, err := vault.Find(findUserURL, cherURL)
	if err != nil {
		fmt.Println("Ничего не найдено")
	}
	fmt.Println(findAccounts)
	color.Green("это все что нашлось")
}

func cherURL(acc account.Account, str string) bool {
	return strings.Contains(acc.Url, str)
}

func promptData[T string | []string](slice T) string {
	switch v := any(slice).(type) {
	case string:
		fmt.Print(v + " ")
		var res string
		fmt.Scan(&res)
		return res
	case []string:
		for i, value := range v {
			if i == len(v)-1 {
				fmt.Print(value + ": ")
			} else {
				fmt.Println(value)
			}
		}
		return ""
	default:
		return ""
	}
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
