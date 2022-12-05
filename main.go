package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/socket", socketHandler)
	//err := Connect("localhost", "5433", "webSocket", "local", "local", SSLModeDisable)
	//if err != nil {
	//	fmt.Printf(" err: %v", err)
	//}
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode("Testing")
		if err != nil {
			return
		}
	})
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal(err)
	}
}
