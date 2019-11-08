package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	
)

func TestCreateArticle(t *testing.T) {

	var jsonStr = []byte(`{"id":"1","title":"Web api using Go lang","date":"2016-09-22","body":"Testing files","tags":["health","science"]}`)

	req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createArticle)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := `{"id":"1","title":"Web api using Go lang","date":"2016-09-22","body":"Testing files","tags":["health","science"]}`
	getResult := strings.TrimSuffix(rr.Body.String(), "\n")
	
	if getResult != expected {
		t.Errorf("handler returned unexpected body")
		t.Errorf("Got : %v", rr.Body.String())
		t.Errorf("Want: %v", expected)
	

	}
}

