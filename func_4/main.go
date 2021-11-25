package func_4

import (
	"fmt"
	"net/http"
)

// Adds two integers
func Add(a int, b int) int {
	var c = a + b
	fmt.Println(c)
	return c
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, friend")
}
