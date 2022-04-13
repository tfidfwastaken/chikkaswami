package main

import (
	"fmt"
	"net/http"

	"chikkaswami/chikkaswami"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", fallback)
	return mux
}

func fallback(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "We don't shorten that here.")
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/chinnaswamy": "https://github.com/nilenso/chinnaswamy",
		"/chikkaswami": "https://github.com/tfidfwastaken/chikkaswami",
	}
	mapHandler := chikkaswami.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")
	err := http.ListenAndServe(":8080", mapHandler)
	if err != nil {
		fmt.Println("Error in starting server.")
		return
	}
}
