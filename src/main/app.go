package main

import (
    "baset/src/logic-parser/ast"
    "baset/src/netUtils"
	"fmt"
	"log"
	"net/http"
)


func main() {
    const port = ":9090"

    tree := ast.Parse("1&1")
    fmt.Println(tree.R())

    http.HandleFunc("/api", server.QueryPost)
    err := http.ListenAndServe(port, nil)
    if err == nil {
        log.Println("Server running on port", port)
    } else {
        log.Fatal(err)
    }

}

