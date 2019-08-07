package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8000", "port")
	dir := flag.String("d", ".", "dir")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Printf("Serving %s on Http port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

// import "fmt"

// func main() {
// 	var input string
// 	fmt.Printf("Enter your Name: ")
// 	fmt.Scanln(&input)
// 	fmt.Printf("Hello, %s!", input)
// }
