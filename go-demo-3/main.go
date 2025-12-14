package main

import "fmt"

func main() {
	// m := map[string]string{
	// 	"Ключ":  "значение",
	// 	"ключ2": "значение 2",
	// }

	// fmt.Println(m)
	// fmt.Println(m["Ключ"])
	// m["yandex"] = "значение яндекса"

	// delete(m, "yandex")
	// fmt.Println(m)

	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	var bookmarks = map[string]string{}
	var comandUser string
	loop:
	for {
		fmt.Println("Нажимите следующее для работы программы")
		fmt.Println("1 - просмотреть доступные закладки")
		fmt.Println("2 - создать новую закладку")
		fmt.Println("3 - удалить закладку по имени")
		fmt.Println("4 - выйти")

		fmt.Scan(&comandUser)
		fmt.Println(bookmarks, comandUser)

		switch comandUser {
		case "1":
			fmt.Println("Функционал просмотра всех закладок")
			if len(bookmarks) == 0 {
				fmt.Println("Список закладок пуст")
			} else {
				for key, val := range bookmarks {
					fmt.Println(key, val)
				}
			}
		case "2":
			fmt.Println("Функционал создания новой закладки")
			var localKey string
			var localValue string
			fmt.Println("введите название")
			fmt.Scan(&localKey)
			fmt.Println("введите ссылку")
			fmt.Scan(&localValue)
			bookmarks[localKey] = localValue
			fmt.Println("заметка успешно добавлена")
		case "3":
			fmt.Println("Функционал удалить закладку по имени")
			var remuveLocalName string
			for key, val := range bookmarks {
				fmt.Println(key, val)
			}
			fmt.Println("введите имя для удаления")
			fmt.Scan(&remuveLocalName)
			delete(bookmarks, remuveLocalName)
			fmt.Println("заметка успешно удалена")
		case "4":
			fmt.Println("выйти")
			break loop
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
