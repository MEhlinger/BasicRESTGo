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
var artifacts []Artifact

func getArtifacts(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(artifacts)
}

func getArtifact(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, artifact := range artifacts {
        if artifact.ID == params["id"] {
            json.NewEncoder(w).Encode(artifact)
            return
        }
    }
    json.NewEncoder(w).Encode(&Artifact{})
}

func createArtifact(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var artifact Artifact
    _ = json.NewDecoder(req.Body).Decode(&artifact)
    artifact.ID = params["id"]
    artifacts = append(artifacts, artifact)
    json.NewEncoder(w).Encode(artifacts)
}

func deleteArtifact(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, artifact := range artifacts {
        if artifact.ID == params["id"] {
            artifacts = append(artifacts[:index], artifacts[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(artifacts)
}

func initData() {
    artifacts = append(artifacts, Artifact{"1", "+2 Boots of Flight", "Provides the user with the power of flight"})
    artifacts = append(artifacts, Artifact{"2", "+7 Sword of Reckoning", "You better believe it, buster."})
}

func main() {
    r := mux.NewRouter()
    initData()
    r.HandleFunc("/artifacts", getArtifacts).Methods("GET")
    r.HandleFunc("/artifacts/{id}", getArtifact).Methods("GET")
    r.HandleFunc("/artifacts/{id}", createArtifact).Methods("POST")
    r.HandleFunc("/artifacts/{id}", deleteArtifact).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12345",r))
}
