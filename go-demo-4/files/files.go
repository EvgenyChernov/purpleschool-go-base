package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile() {
	color.Cyan("Функционал чтения файла")
}

func WriteFile(content string, name string) {
	color.Green("Функционал записи файла")
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		fmt.Println(err)
	}
	fmt.Println("Запись успешна")
}
