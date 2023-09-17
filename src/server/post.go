package server

import (
	"encoding/json"
	"log"
	"net/http"
	"setbase/src/db"
	"strings"
)

var database *db.DB

func EnableEndpoint(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Methods", "\"POST\", \"OPTIONS\"");
    (*w).Header().Set("Access-Control-Allow-Origin", "*");
    (*w).Header().Set("Access-Control-Allow-Headers", "*");
}

func GeneralHandle(w http.ResponseWriter, r *http.Request) {
    if ((*r).Method == "OPTIONS") {
        EnableEndpoint(&w);
        w.WriteHeader(200);
    } else {
        w.Write([]byte("hola " + r.RemoteAddr));
    }
}
func QueryPost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    EnableEndpoint(&w)

    var data db.Query
    decoder := json.NewDecoder(r.Body)
    decodeErr := decoder.Decode(&data)
    if decodeErr != nil {
        log.Println(decodeErr)
    }

    var startErr error

    if database == nil {
        database, startErr = db.StartDb()
        if startErr != nil {
            log.Println("db start error: " + startErr.Error())
            w.WriteHeader(http.StatusInternalServerError)
        }
    }

    queryResult, queryErr := database.Query(data)
    if queryErr != nil {
        log.Println("query result error: " + queryErr.Error())
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    response := new(strings.Builder)
    encodeErr := json.NewEncoder(response).Encode(queryResult)
    if encodeErr != nil {
        log.Println("json encode error: " + encodeErr.Error())
        w.Write([]byte(response.String()))
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Write([]byte(response.String()))
    return
}


