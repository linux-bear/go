package goadd

import "fmt"

func Addme(num1 int, num2 int) int {
	returned := num1 + num2
	fmt.Println(returned)
	return returned
}
