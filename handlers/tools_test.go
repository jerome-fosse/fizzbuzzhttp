package handlers

import (
	"net/http"
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
	params := parseQueryParameters(r.URL)

	// Then the parameters are equall to the default values
	assert.Equal(t, 3, params.int1)
	assert.Equal(t, "fizz", params.word1)
	assert.Equal(t, 5, params.int2)
	assert.Equal(t, "buzz", params.word2)
	assert.Equal(t, 15, params.limit)
}

func Test_should_return_parameters_values_when_fizzbuzz_parameters_are_in_the_url(t *testing.T) {
	// Given a fizzbuzz request with parameters
	r, err := http.NewRequest(http.MethodGet, "/fizzbuzz?int1=4&word1=titi&int2=7&word2=toto&limit=30", nil)
	if err != nil {
		t.Fatal(err)
	}

	// When I parse the URL to find the parameters
	params := parseQueryParameters(r.URL)

	// Then the parameters are equall to the default values
	assert.Equal(t, 4, params.int1)
	assert.Equal(t, "titi", params.word1)
	assert.Equal(t, 7, params.int2)
	assert.Equal(t, "toto", params.word2)
	assert.Equal(t, 30, params.limit)
}

func Test_should_return_dafault_value_when_a_numeric_parameter_value_is_nan(t *testing.T) {
	// Given a fizzbuzz request with parameters
	r, err := http.NewRequest(http.MethodGet, "/fizzbuzz?int1=ab&word1=titi&int2=cd&word2=toto&limit=ef", nil)
	if err != nil {
		t.Fatal(err)
	}

	// When I parse the URL to find the parameters
	params := parseQueryParameters(r.URL)

	// Then the parameters are equall to the default values
	assert.Equal(t, 3, params.int1)
	assert.Equal(t, "titi", params.word1)
	assert.Equal(t, 5, params.int2)
	assert.Equal(t, "toto", params.word2)
	assert.Equal(t, 15, params.limit)
}
