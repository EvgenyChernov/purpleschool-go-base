package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	login    string `json:"login" xml:"test"`
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.login, acc.password, acc.url)
	// fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) GeneratPassword() {
	var lenghtPassword int = 10
	var localPassword string
	for i := 0; i < lenghtPassword; i++ {
		localPassword += randomLetter()
	}
	acc.password = localPassword
}

func NewAccountWithTimeStam(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" || password == "" || urlString == "" {
		return nil, errors.New("invalid input")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStamp{
		Account: Account{
			url:      urlString,
			password: password,
			login:    login,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	fmt.Println(string(field.Tag))
	return newAcc, nil
}

func randomLetter() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return string(letters[rand.Intn(len(letters))])
}
