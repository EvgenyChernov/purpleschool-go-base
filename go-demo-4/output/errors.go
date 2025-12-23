package output

import "fmt"

func PrintError(value any) {

	switch value.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
	fmt.Println(value)
}
