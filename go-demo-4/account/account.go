package account

import (
	"errors"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login" xml:"test"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.Login, acc.Password, acc.Url)
	// fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) GeneratPassword() {
	var lenghtPassword int = 10
	var localPassword string
	for i := 0; i < lenghtPassword; i++ {
		localPassword += randomLetter()
	}
	acc.Password = localPassword
}

func NewAccount(Login, Password, UrlString string) (*Account, error) {
	if Login == "" || Password == "" || UrlString == "" {
		return nil, errors.New("invalid input")
	}

	_, err := url.ParseRequestURI(UrlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		Url:       UrlString,
		Password:  Password,
		Login:     Login,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	// fmt.Println(string(field.Tag))
	return newAcc, nil
}

func randomLetter() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return string(letters[rand.Intn(len(letters))])
}
