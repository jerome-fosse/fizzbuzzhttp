package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_return_default_value_when_fizzbuzz_parameter_are_not_found(t *testing.T) {
	// Given a fizzbuzz request without parameters
	r, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	// When I parse the URL to find the parameters
	int1, word1, int2, word2, limit := parseRequestParameters(r.URL)

	// Then the parameters are equall to the default values
	assert.Equal(t, 3, int1)
	assert.Equal(t, "fizz", word1)
	assert.Equal(t, 5, int2)
	assert.Equal(t, "buzz", word2)
	assert.Equal(t, 15, limit)
}

func Test_should_return_parameters_values_when_fizzbuzz_parameters_are_in_the_url(t *testing.T) {
	// Given a fizzbuzz request with parameters
	r, err := http.NewRequest(http.MethodGet, "/fizzbuzz?int1=4&word1=titi&int2=7&word2=toto&limit=30", nil)
	if err != nil {
		t.Fatal(err)
	}

	// When I parse the URL to find the parameters
	int1, word1, int2, word2, limit := parseRequestParameters(r.URL)

	// Then the parameters are equall to the default values
	assert.Equal(t, 4, int1)
	assert.Equal(t, "titi", word1)
	assert.Equal(t, 7, int2)
	assert.Equal(t, "toto", word2)
	assert.Equal(t, 30, limit)
}

func Test_should_return_dafault_value_when_a_numeric_parameter_value_is_nan(t *testing.T) {
	// Given a fizzbuzz request with parameters
	r, err := http.NewRequest(http.MethodGet, "/fizzbuzz?int1=ab&word1=titi&int2=cd&word2=toto&limit=ef", nil)
	if err != nil {
		t.Fatal(err)
	}

	// When I parse the URL to find the parameters
	int1, word1, int2, word2, limit := parseRequestParameters(r.URL)

	// Then the parameters are equall to the default values
	assert.Equal(t, 3, int1)
	assert.Equal(t, "titi", word1)
	assert.Equal(t, 5, int2)
	assert.Equal(t, "toto", word2)
	assert.Equal(t, 15, limit)
}

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

	// Then I have a method not allowed error (Http Satatus Code = 405
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

	// Then I have a method not allowed error (Http Satatus Code = 405
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

	// Then I have a method not allowed error (Http Satatus Code = 405
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

	// Then I have a method not allowed error (Http Satatus Code = 405
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

	// Then I have a method not allowed error (Http Satatus Code = 405
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
	assert.Equal(t, "Method OPTION not allowed.", rec.Body.String())
}

func Test_Get_Fizzbuzz_without_parameters_should_return_a_list_of_15_elements(t *testing.T) {
	// Given a PUT on /fizzbuzz
	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	// When I execute the request
	handler.ServeHTTP(rec, req)

	// Then I have a fizzbuzz result whith 15 elements
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"result\":[\"1\",\"2\",\"fizz\",\"4\",\"buzz\",\"fizz\",\"7\",\"8\",\"fizz\",\"buzz\",\"11\",\"fizz\",\"13\",\"14\",\"fizzbuzz\"]}", rec.Body.String())
}

func Test_Get_Fizzbuzz_without_parameters_should_be_ok(t *testing.T) {
	// Given a PUT on /fizzbuzz
	req, err := http.NewRequest(http.MethodGet, "/fizzbuzz?limit=20&word1=titi&word2=toto&int1=4&int2=7", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(FizzbuzzHandler)

	// When I execute the request
	handler.ServeHTTP(rec, req)

	// Then I have a fizzbuzz result
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"result\":[\"1\",\"2\",\"3\",\"titi\",\"5\",\"6\",\"toto\",\"titi\",\"9\",\"10\",\"11\",\"titi\",\"13\",\"toto\",\"15\",\"titi\",\"17\",\"18\",\"19\",\"titi\"]}", rec.Body.String())
}