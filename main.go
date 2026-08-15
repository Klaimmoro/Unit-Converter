package main

import (
	"encoding/json"
	"log"
	"net/http"
	"unit-converter/src"
	"unit-converter/src/kernel"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req kernel.ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	var unit_converter src.UnitConverter
	result, err := unit_converter.Convert(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(kernel.ConvertResponse{Result: result})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/convert", handleFunc)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./src/html/page.html")
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error: %s\n", err)
	}

}
