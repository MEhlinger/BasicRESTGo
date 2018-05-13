package main

import (
    "net/http"
    "encoding/json"
    "log"

    "github.com/gorilla/mux"
)

type Artifact struct {
    ID      string  `json:"id,omitempty"`
    Name    string  `json:"name,omitempty"`
    Detail  string  `json:"detail,omitempty"`
}

// In lieu of proper DB connection, this is our data source of record
var items []Artifact

func getArtifacts(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(items)
}

func getArtifact(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range items {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Artifact{})
}

func createArtifact(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var item Artifact
    _ = json.NewDecoder(req.Body).Decode(&item)
    item.ID = params["id"]
    items = append(items, item)
    json.NewEncoder(w).Encode(items)
}

func initData() {
    items = append(items, Artifact{"1", "+2 Boots of Flight", "Provides the user with the power of flight"})
    items = append(items, Artifact{"2", "+7 Sword of Reckoning", "You better believe it, buster."})
}

func main() {
    r := mux.NewRouter()
    initData()
    r.HandleFunc("/artifacts", getArtifacts).Methods("GET")
    r.HandleFunc("/artifacts/{id}", getArtifact).Methods("GET")
    r.HandleFunc("/artifacts/{id}", createArtifact).Methods("POST")
    log.Fatal(http.ListenAndServe(":12345",r))
}
