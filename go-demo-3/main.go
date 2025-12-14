package main

import "fmt"

func main() {
	m := map[string]string{
		"Ключ": "значение",
		"ключ2": "значение 2",
	}

	fmt.Println(m)
	fmt.Println(m["Ключ"])
	m["yandex"] = "значение яндекса"

	delete(m, "yandex")
	fmt.Println(m)

    for key, value := range m{
		fmt.Println(key, value)
	}

}
