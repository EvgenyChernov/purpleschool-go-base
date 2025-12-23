package output

import "fmt"

func PrintError(value any) {
	value, ok := value.(int)
	if ok {
		fmt.Println("int")
		return
	}
	stringValue, ok := value.(string)
	if ok {
		fmt.Println("string")
		return
	}
	errorValue, ok := value.(error)
	if ok {
		fmt.Println(errorValue.Error())
		return
	}

	switch value.(type) {
	case int:
		fmt.Println(value)
	case string:
		fmt.Println(stringValue)
	default:
		fmt.Println("unknown")
	}
	fmt.Println(value)
}
