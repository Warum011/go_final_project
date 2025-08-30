package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Warum011/go_final_project/pkg/api"
)

func StartServer(webDir string) {

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	addr := fmt.Sprintf(":%s", port)

	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", fs)

	log.Printf("Starting server on http://localhost%s/, serving %s\n", addr, webDir)

	api.Init()

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
