package func_4

import (
	"fmt"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, friend")
}
