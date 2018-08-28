package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("HOSTNAME")
	if hostname != "" {
	  fmt.Fprintf(w,hostname)
        } else {
	  fmt.Fprintf(w,"HOSTNAME is empty")
        }
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
