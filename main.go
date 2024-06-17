package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// ServeHandler function to serve the policy.html file
func ServeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "policy.html")
}

func main() {
	fmt.Println("starting your http server!")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//Code to serve the policy.html file
	http.HandleFunc("/", ServeHandler)
	http.ListenAndServe(":"+port, nil)
}
