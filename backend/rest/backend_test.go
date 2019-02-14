package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerGood(t *testing.T) {

	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle(func(w http.ResponseWriter, req *http.Request) (interface{}, error) { return []string{"test"}, nil }))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "wrong status")

	assert.Equal(t, rr.Body.String(), `{"success":true,"data":["test"]}`)
}

func TestHandlerBad(t *testing.T) {

	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle(func(w http.ResponseWriter, req *http.Request) (interface{}, error) { return nil, errors.New("error") }))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "wrong status")

	assert.Equal(t, rr.Body.String(), `{"success":false,"message":"Cannot get data"}`)
}