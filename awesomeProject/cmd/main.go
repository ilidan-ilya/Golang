package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nowhere *int
	var i, z, err int
	var a, b string
	fmt.Println("Введите переменную а:")
	_, err = fmt.Scanln("%s", &a)
	if i != nil {
		fmt.Println(i)
	}

	fmt.Println("Введите переменную b:")
	fmt.Scanln("%s\n", &b)
	i = strconv.Atoi(a)
	z = strconv.Atoi(b)

}
