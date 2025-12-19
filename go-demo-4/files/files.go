package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile() {
	file, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("error read file")
		return
	}
	fmt.Println(string(file))
}

func WriteFile(content []byte, name string) {
	color.Green("Функционал записи файла")
	file, err := os.Create(name)
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
