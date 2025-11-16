package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello world</h1>")
	})

	http.ListenAndServe(":"+PORT, nil)
}
