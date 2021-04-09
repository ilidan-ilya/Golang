package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})
	http.ListenAndServe(":8000", nil)

	var a, b string
	fmt.Println("Введите переменную а:")
	_, gut := fmt.Scanln("%s", &a)
	if a != nil {
		fmt.Println(a)
	}

	fmt.Println("Введите переменную b:")
	fmt.Scanln("%s\n", &b)
	i := strconv.Atoi(a)
	z := strconv.Atoi(b)
	fmt.Println("Сложение ", a, "+", b)
}
