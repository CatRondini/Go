package main

import (
    "fmt"
    "net/http"
    "errors"
    "os"
)

func helloWorldPage(writer http.ResponseWriter, req *http.Request) {
    fs := http.FileServer(http.Dir("Pages/"))
    http.Handle("/index.html/", http.StripPrefix("/index.html/", fs))
}

func main() {
    http.HandleFunc("/", helloWorldPage)
    
    err := http.ListenAndServe(":3333", nil)
    
    
    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
    }

}