package main

import (
	"net/http"
	"os"
	
	"github.com/zhuharev/ssr"
)

func main() {
	renderer := ssr.New()
	http.Handle("/render",renderer)
	port := "5000"
	if os.Getenv("PORT")!="" {
		port = os.Getenv("PORT")
	}
    http.ListenAndServe(":"+port, nil)
}