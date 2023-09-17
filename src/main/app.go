package main

import (
	"setbase/src/db"
	"setbase/src/server"
	"log"
	"net/http"

	// test data package
)

func main() {
    const port string = "9090"

    //dd := &db.DB{}
    //dd.Fill(test_db.DB)

    db.StartDb()

    http.HandleFunc("/", server.GeneralHandle)
    http.HandleFunc("/media/", server.StaticServe)
    http.HandleFunc("/api", server.QueryPost)
    http.HandleFunc("/api/", server.QueryPost)

    log.Println("Opening server @" + port)
    serverErr := http.ListenAndServe(":" + port, nil)
    if serverErr != nil {
        log.Println(serverErr)
    }
}

