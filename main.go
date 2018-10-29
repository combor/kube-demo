package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	version := "1.0"
	hostname := os.Getenv("HOSTNAME")
	fmt.Fprintf(w,"hostname: %s\nversion: %s\n", hostname, version)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("No PORT variable set. Exiting...")
		return
	}
	fmt.Printf("Listening on 0.0.0.0:%s\n", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
