package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4}

	reverse(a)

	fmt.Println(a)
}

func reverse(a []int) {
	for i := 0; i < len(a)/2; i++ {
		tmp := a[i]
		a[i] = a[len(a)-1-i]
		a[len(a)-1-i] = tmp
	}
}
