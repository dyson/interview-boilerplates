package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":80"
	if appAddr := os.Getenv("APP_ADDR"); appAddr != "" {
		addr = appAddr
	}

	log.Println("listening on", addr)
	if err := http.ListenAndServe(addr, handler()); err != nil {
		log.Fatalln("listen and serve error:", err)
	}
}

func handler() http.Handler {
	http.HandleFunc("/hello", hello)
	return http.DefaultServeMux
}

func hello(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(map[string]string{"hello": "world"})
	if err != nil {
		log.Println("handler hello can't marshal json:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(body)
}
