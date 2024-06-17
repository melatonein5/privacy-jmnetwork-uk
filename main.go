package main

import (
	"net/http"
)

// ServeHandler function to serve the policy.html file
func ServeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "policy.html")
}

func main() {
	//Code to serve the policy.html file
	http.HandleFunc("/", ServeHandler)
	http.ListenAndServe(":80", nil)
}
