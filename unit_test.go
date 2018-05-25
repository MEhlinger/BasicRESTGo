package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "bytes"
    "os"
    "strings"

    "github.com/gorilla/mux"
)

var  testArtifact = []byte(`{"name":"testArtifact","detail":"This is a test artifact."}`)

func TestGetArtifacts(t *testing.T) {
    expected := `[{"id":"1","name":"+2 Boots of Flight","detail":"Provides the user with the power of flight"},{"id":"2","name":"+7 Sword of Reckoning","detail":"You better believe it, buster."}]`
    req, _ := http.NewRequest("GET", "/artifacts", nil)
    recorder := httptest.NewRecorder()
    getArtifacts(recorder, req)
    if actual := strings.TrimSpace(recorder.Body.String()); actual != expected {
        t.Errorf("Response was:\n%v\nexpected:\n%v\n", actual, expected)
    }
}

func TestGetArtifact(t *testing.T) {
    expected := `{"id":"2","name":"+7 Sword of Reckoning","detail":"You better believe it, buster."}`
    req, _ := http.NewRequest("GET", "artifacts/2", nil)
    req = mux.SetURLVars(req, map[string]string{"id":"2"})
    recorder := httptest.NewRecorder()
    getArtifact(recorder, req)
    if actual := strings.TrimSpace(recorder.Body.String()); actual != expected {
        t.Errorf("Response was:\n%v\nexpected:\n%v\n", actual, expected)
    }
}

func TestCreateArtifact(t *testing.T) {
    expected := `[{"id":"1","name":"+2 Boots of Flight","detail":"Provides the user with the power of flight"},{"id":"2","name":"+7 Sword of Reckoning","detail":"You better believe it, buster."},{"id":"3","name":"testArtifact","detail":"This is a test artifact."}]`
    req, _ := http.NewRequest("POST", "/artifacts/3", bytes.NewBuffer(testArtifact))
    req = mux.SetURLVars(req, map[string]string{"id":"3"})
    recorder := httptest.NewRecorder()
    createArtifact(recorder, req)
    if actual := strings.TrimSpace(recorder.Body.String()); actual != expected {
        t.Errorf("Response was:\n%v\nexpected:\n%v\n", actual, expected)
    }
}

func TestCreateArtifactDuplicateID(t *testing.T) {
    req, _ := http.NewRequest("POST", "/artifacts/3", bytes.NewBuffer(testArtifact))
    req = mux.SetURLVars(req, map[string]string{"id":"3"})
    recorder := httptest.NewRecorder()
    createArtifact(recorder, req)
    if recorder.Code != http.StatusBadRequest {
        t.Errorf("Status code returned was %v, expected 400", recorder.Code)
    }
}

func TestDeleteArtifact(t *testing.T) {
    expected := `[{"id":"1","name":"+2 Boots of Flight","detail":"Provides the user with the power of flight"},{"id":"2","name":"+7 Sword of Reckoning","detail":"You better believe it, buster."}]`
    req, _ := http.NewRequest("DELETE", "/artifacts/3", nil)
    req = mux.SetURLVars(req, map[string]string{"id":"3"})
    recorder := httptest.NewRecorder()
    deleteArtifact(recorder, req)
    if actual := strings.TrimSpace(recorder.Body.String()); actual != expected {
        t.Errorf("Response was:\n%v\nexpected:\n%v\n", actual, expected)
    }
}

func setup() {
    initData()
}

func TestMain(m *testing.M) {
    setup()
    runTests := m.Run()
    os.Exit(runTests)
}


