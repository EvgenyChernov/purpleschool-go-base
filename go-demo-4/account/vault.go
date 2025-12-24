package account

import (
	"demo/password/encrypter"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write(content []byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault     Vault
	Db        Db
	Encrypter *encrypter.Encrypter
}

func (vault *Vault) FindByURLtoDelete(findUserURL string) (int, error) {
	var newAccounts []Account
	deletedCount := 0

	for _, acc := range vault.Accounts {
		if acc.Url == findUserURL {
			deletedCount++
			continue // пропускаем элемент, который нужно удалить
		}
		newAccounts = append(newAccounts, acc)
	}

	vault.Accounts = newAccounts // обновляем срез

	if deletedCount == 0 {
		return 0, fmt.Errorf("no accounts found for URL: %s", findUserURL)
	}

	return deletedCount, nil
}

func (vault *VaultWithDb) FindToURL(findUserURL string) ([]Account, error) {
	var result []Account

	for _, acc := range vault.Vault.Accounts {
		if acc.Url == findUserURL {
			result = append(result, acc)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no accounts found for url: %s", findUserURL)
	}

	return result, nil
}

func (vault *VaultWithDb) Find(str string, checker func(Account, string) bool) ([]Account, error) {
	var result []Account

	for _, acc := range vault.Vault.Accounts {
		isMatched := checker(acc, str)
		if isMatched {
			result = append(result, acc)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no accounts found for: %s", str)
	}

	return result, nil
}

func NewVault(db Db, encrypter *encrypter.Encrypter) *VaultWithDb {

	encryptedData, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Db: db,
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			Encrypter: encrypter,
		}
	}

	// Проверяем, что файл не пустой
	if len(encryptedData) == 0 {
		return &VaultWithDb{
			Db: db,
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			Encrypter: encrypter,
		}
	}

	// Расшифровываем данные
	var file []byte
	if encrypter != nil {
		file, err = encrypter.Decrypt(encryptedData)
		if err != nil {
			fmt.Println("ошибка расшифровки файла:", err)
			return &VaultWithDb{
				Db: db,
				Vault: Vault{
					Accounts:  []Account{},
					UpdatedAt: time.Now(),
				},
				Encrypter: encrypter,
			}
		}
	} else {
		file = encryptedData
	}

	// Проверяем, что после расшифровки данные не пустые
	if len(file) == 0 {
		return &VaultWithDb{
			Db: db,
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			Encrypter: encrypter,
		}
	}

	color.Blue("МАЙ файл успешно прочитан")

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println("ошибка чтения файла:", err)
		return &VaultWithDb{
			Db: db,
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			Encrypter: encrypter,
		}
	}
	return &VaultWithDb{
		Vault:     vault,
		Db:        db,
		Encrypter: encrypter,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Vault.Accounts = append(vault.Vault.Accounts, acc)
	vault.Vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		fmt.Println("Ошибка преобразования")
		return
	}

	// Шифруем данные перед записью
	var encryptedData []byte
	if vault.Encrypter != nil {
		encryptedData, err = vault.Encrypter.Encrypt(data)
		if err != nil {
			fmt.Println("Ошибка шифрования:", err)
			return
		}
	} else {
		encryptedData = data
	}

	vault.Db.Write(encryptedData)
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
