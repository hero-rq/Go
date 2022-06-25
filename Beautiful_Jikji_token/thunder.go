package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(rw, "panda")
}

func main() {
  http.HandleFunc("/", home)
  fmt.Printf("listening on server http://local%s", port)
  log.Fatal(http.ListenAndServe(port, nil))
}
