// +build integration
// Integration Tests

package itest

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz_OK(t *testing.T) {
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "{\"result\":[\"1\",\"2\",\"fizz\",\"4\",\"buzz\",\"fizz\",\"7\",\"8\",\"fizz\",\"buzz\",\"11\",\"fizz\",\"13\",\"14\",\"fizzbuzz\"]}", string(bytes))
}

func Test_FizzBuzz_WithParams_OK(t *testing.T) {
	client := &http.Client{}
	r, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz?int1=4&word1=Hello&int2=6&word2=World&limit30", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "{\"result\":[\"1\",\"2\",\"3\",\"Hello\",\"5\",\"World\",\"7\",\"Hello\",\"9\",\"10\",\"11\",\"HelloWorld\",\"13\",\"14\",\"15\"]}", string(bytes))
}
