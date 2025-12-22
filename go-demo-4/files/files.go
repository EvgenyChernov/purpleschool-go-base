package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {

	return &JsonDb{
		fileName: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	file, err := os.ReadFile(db.fileName)
	if err != nil {
		fmt.Println("error read file")
		return nil, err
	}
	return file, nil
	// fmt.Println(string(file))
}

func (db *JsonDb) Write(content []byte) {
	color.Green("Функционал записи файла")
	file, err := os.Create(db.fileName)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	if err != nil {
		defer file.Close()
		fmt.Println(err)
	}
	fmt.Println("Запись успешна")
}
