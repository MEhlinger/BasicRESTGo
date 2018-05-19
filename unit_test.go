package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "bytes"
)

var  testArtifact = []byte(`{"name":"testArtifact","detail":"This is a test artifact."}`)

func TestGetArtifacts(t *testing.T) {
    req, _ := http.NewRequest("GET", "/artifacts", nil)
    recorder := httptest.NewRecorder()
    getArtifacts(recorder, req)
    if recorder.Code != 200 {
        t.Errorf("Status code returned was %v , expected 200.", recorder.Code)
    }
}

func TestGetArtifact(t *testing.T) {
    req, _ := http.NewRequest("GET", "/artifacts/1", nil)
    recorder := httptest.NewRecorder()
    getArtifact(recorder, req)
    if recorder.Code != 200 {
        t.Errorf("Status code returned was %v, expected 200.", recorder.Code)
    }
}

func TestCreateArtifact(t *testing.T) {
    req, _ := http.NewRequest("POST", "/artifacts/3", bytes.NewBuffer(testArtifact))
    recorder := httptest.NewRecorder()
    createArtifact(recorder, req)
    if recorder.Code != 200 {
        t.Errorf("Status code returned was %v , expected 200.", recorder.Code)
    }
}


func TestCreateArtifactDuplicateID(t *testing.T) {
    req, _ := http.NewRequest("POST", "/artifacts/1", bytes.NewBuffer(testArtifact))
    recorder := httptest.NewRecorder()
    createArtifact(recorder, req)
    if recorder.Code != 500 {
        t.Errorf("Status code returned was %v, expected 500", recorder.Code)
    }
}

