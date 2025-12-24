package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// fmt.Println(os.Getenv("VAR"))

	output.PrintError(1)
	output.PrintError("1")
	// files.WriteFile("Привет Мир! я ФАЙЛ !!!!", "file.txt")
	// files.ReadFile()
	encrypter, err := encrypter.NewEncrypter()
	if err != nil {
		fmt.Println("Error creating encrypter")
	}
	vault := account.NewVault(files.NewJsonDb("data.json"), encrypter)
	// vault := account.NewVault(cloud.NewCloudDb("https://cloud.json"))
	color.Green("1. Создать аккаунт")
	color.Blue("2. Найти аккаунт по URL")
	color.Green("3. Найти аккаунт по логину")
	color.Yellow("4. Уадалить аккаунт")
	color.Red("5. Выход")
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

func findAccountByURL(vault *account.VaultWithDb) {
	color.Blue("Введите URL")
	var findUserURL string
	fmt.Scan(&findUserURL)
	findAccounts, err := vault.Find(findUserURL, func(a account.Account, s string) bool {
		return strings.Contains(strings.ToLower(a.Url), strings.ToLower(s))
	})
	if err != nil {
		fmt.Println("Ничего не найдено")
		return
	}
	fmt.Println(findAccounts)
	color.Green("это все что нашлось")
}

func findAccountByLogin(vault *account.VaultWithDb) {
	color.Green("Введите логин")
	var findUserLogin string
	fmt.Scan(&findUserLogin)
	findAccounts, err := vault.Find(findUserLogin, func(a account.Account, s string) bool {
		return strings.Contains(strings.ToLower(a.Login), strings.ToLower(s))
	})
	if err != nil {
		fmt.Println("Ничего не найдено")
		return
	}
	fmt.Println(findAccounts)
	color.Green("это все что нашлось")
}

func promptData[T string | []string](slice ...T) string {
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
