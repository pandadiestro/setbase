package server

import (
	"log"
	"net/http"
	"os"
)

func StaticServe(w http.ResponseWriter, r *http.Request) {
    EnableEndpoint(&w)
    w.Header().Set("Content-Type", "image/png")

    if r.Method != http.MethodGet && r.Method != http.MethodOptions {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    path := r.URL.Query().Get("path")
    log.Println(path)
    bytesFile, pathErr := os.ReadFile(path)

    if pathErr != nil {
        w.WriteHeader(http.StatusNotFound)
    }

    w.Write(bytesFile)
}



