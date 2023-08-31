package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {

	var defaultPath string
	flag.StringVar(&defaultPath, "path", "public", "The default path for hosting")
	flag.Parse()

	fmt.Println("Default Path:", defaultPath)

	fileServer := http.FileServer(http.Dir(defaultPath))
	http.Handle("/", http.StripPrefix("/", fileServer))
	http.ListenAndServe(":8080", nil)
}
