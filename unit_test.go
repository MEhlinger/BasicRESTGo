package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "bytes"
    "fmt"
    "os"

    "github.com/gorilla/mux"
)

var  testArtifact = []byte(`{"name":"testArtifact","detail":"This is a test artifact."}`)

func TestGetArtifacts(t *testing.T) {
    req, _ := http.NewRequest("GET", "/artifacts", nil)
    recorder := httptest.NewRecorder()
    getArtifacts(recorder, req)
    fmt.Printf("\ngetArtifacts:\n%v\n", recorder.Body)
    if recorder.Code != http.StatusOK {
        t.Errorf("Status code returned was %v , expected http.StatusOK.", recorder.Code)
    }
}

func TestGetArtifact(t *testing.T) {
    req, _ := http.NewRequest("GET", "artifacts/2", nil)
    req = mux.SetURLVars(req, map[string]string{"id":"2"})
    recorder := httptest.NewRecorder()
    getArtifact(recorder, req)
    fmt.Printf("\ngetArtifact:\n%v\n", recorder.Body)
    if recorder.Code != http.StatusOK {
        t.Errorf("Status code returned was %v, expected http.StatusOK.", recorder.Code)
    }
}

func TestCreateArtifact(t *testing.T) {
    req, _ := http.NewRequest("POST", "/artifacts/3", bytes.NewBuffer(testArtifact))
    req = mux.SetURLVars(req, map[string]string{"id":"3"})
    recorder := httptest.NewRecorder()
    createArtifact(recorder, req)
    fmt.Printf("\ncreateArtifact:\n%v\n", recorder.Body)
    if recorder.Code != http.StatusOK {
        t.Errorf("Status code returned was %v , expected http.StatusOK.", recorder.Code)
    }
}

func TestCreateArtifactDuplicateID(t *testing.T) {
    req, _ := http.NewRequest("POST", "/artifacts/3", bytes.NewBuffer(testArtifact))
    req = mux.SetURLVars(req, map[string]string{"id":"3"})
    recorder := httptest.NewRecorder()
    createArtifact(recorder, req)
    fmt.Printf("\ncreateArtifact duplicate:\n%v\n", recorder.Body)
    if recorder.Code != http.StatusBadRequest {
        t.Errorf("Status code returned was %v, expected 400", recorder.Code)
    }
}

func TestDeleteArtifact(t *testing.T) {
    req, _ := http.NewRequest("DELETE", "/artifacts/3", nil)
    req = mux.SetURLVars(req, map[string]string{"id":"3"})
    recorder := httptest.NewRecorder()
    deleteArtifact(recorder, req)
    fmt.Printf("\ndeleteArtifact:\n%v\n", recorder.Body)
    if recorder.Code != http.StatusOK {
        t.Errorf("Status code from delete attempt was %v, expected http.StatusOK", recorder.Code)
    }
}

func setup() {
    initData()
    fmt.Printf("\nInitial Data:\n%v\n", artifacts)
}

func TestMain(m *testing.M) {
    setup()
    runTests := m.Run()
    os.Exit(runTests)
}


