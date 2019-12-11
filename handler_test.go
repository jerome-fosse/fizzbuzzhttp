package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_when_POST_method_is_used_on_fizzbuzz_endpoint(t *testing.T) {
	// Given a POST on /fizzbuzz
	req, err := http.NewRequest("POST", "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	// When I execute the request
	handler.ServeHTTP(rec, req)

	// Then I have a method not alloed error (Http Satatus Code = 405
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "Method POST not allowed.", rec.Body.String())
}

func Test_Error_when_PUT_method_is_used_on_fizzbuzz_endpoint(t *testing.T) {
	// Given a PUT on /fizzbuzz
	req, err := http.NewRequest("PUT", "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	// When I execute the request
	handler.ServeHTTP(rec, req)

	// Then I have a method not alloed error (Http Satatus Code = 405
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "Method PUT not allowed.", rec.Body.String())
}

func Test_Error_when_DELETE_method_is_used_on_fizzbuzz_endpoint(t *testing.T) {
	// Given a DELETE on /fizzbuzz
	req, err := http.NewRequest("DELETE", "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	// When I execute the request
	handler.ServeHTTP(rec, req)

	// Then I have a method not alloed error (Http Satatus Code = 405
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "Method DELETE not allowed.", rec.Body.String())
}

func Test_Error_when_PATCH_method_is_used_on_fizzbuzz_endpoint(t *testing.T) {
	// Given a PATCH on /fizzbuzz
	req, err := http.NewRequest("PATCH", "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	// When I execute the request
	handler.ServeHTTP(rec, req)

	// Then I have a method not alloed error (Http Satatus Code = 405
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "Method PATCH not allowed.", rec.Body.String())
}

func Test_Error_when_OPTION_method_is_used_on_fizzbuzz_endpoint(t *testing.T) {
	// Given a PUT on /fizzbuzz
	req, err := http.NewRequest("OPTION", "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	// When I execute the request
	handler.ServeHTTP(rec, req)

	// Then I have a method not alloed error (Http Satatus Code = 405
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "Method OPTION not allowed.", rec.Body.String())
}


