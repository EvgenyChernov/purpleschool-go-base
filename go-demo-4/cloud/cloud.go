package cloud

import (
	"fmt"
	"os"
)

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	file, err := os.ReadFile(db.url)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (db *CloudDb) Write(content []byte) {
	file, err := os.Create(db.url)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	if err != nil {
		defer file.Close()
		fmt.Println(err)
	}
}
