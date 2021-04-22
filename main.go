package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(".env")

	if !os.IsNotExist(err) {
		err := godotenv.Load(".env")

		if err != nil {
			fmt.Println("Error while reading the env file", err)
			panic(err)
		}
	}

	title := os.Getenv("TITLE")
	fmt.Fprintf(w, title)

}

func main() {
	http.HandleFunc("/api", handleFunc)

	log.Println("Server is running:", 3001)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
