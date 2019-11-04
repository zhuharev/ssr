package ssr

import (
	"log"
	"fmt"
	"net/http"
)

func (r *Renderer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("render %s",req.URL)
	targetURL := req.URL.Query().Get("url")
	renderedHTML, err := r.Render(targetURL)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"%s"}`, err)
		return
	}
	w.WriteHeader(200)
	w.Write(renderedHTML)
}
