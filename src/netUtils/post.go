package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type set struct {
    SetName string
    IsInside bool
}

type Query struct {
    Id int
    Sets []set
}

func QueryPost(w http.ResponseWriter, r *http.Request) {
    var queryStruct Query

    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&queryStruct)

    fmt.Println(queryStruct)
}


