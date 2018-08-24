package main

import (
	"fmt"
	"net/http"
	"os"

	figure "github.com/common-nighthawk/go-figure"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname := os.Getenv("HOSTNAME")
	f := figure.NewFigure(hostname, "", true)
	f.Print()
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
