package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	// create request to be sent to handler. In this test we dont have query parameters, so nil is passed.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a response recorder to record the response
	// this satisfies the responsewriter
	rr := httptest.NewRecorder()

	// handler to be tested
	handler := http.HandlerFunc(Home)

	// this handler satifies http.Handler, so we can call serve.http method
	// directly and pass in our request and response recorded
	handler.ServeHTTP(rr, req)

	// confirm the status code is what we want
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code")
	}

	// check if the response body is correct
	expected := `{"Status": "active"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body")
	}

}
