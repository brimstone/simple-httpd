package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		me, _ := os.Hostname()
		fmt.Fprintf(w, "Hello from %s", me)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
