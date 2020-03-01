// +build integration
// Integration Tests

package itest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/object-it/fizzbuzzhttp/statistics"
	"github.com/stretchr/testify/assert"
)

func Test_it_should_display_all_statistics(t *testing.T) {
	client := &http.Client{}

	// I save statistics for several queries
	r1, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		resp, err := client.Do(r1)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	r2, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz?word1=hello&word2=world", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 12; i++ {
		resp, err := client.Do(r2)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	r3, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz?limit=30", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		resp, err := client.Do(r3)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	r4, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz?word1=hello&word2=world&limit=30", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 7; i++ {
		resp, err := client.Do(r4)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	// I request the statistics without a limit
	r, err := http.NewRequest(http.MethodGet, "http://localhost:8080/stats", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// I read the response and unmarshal it
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var stats []statistics.StatisticEntry
	err = json.Unmarshal(bytes, &stats)
	if err != nil {
		t.Fatal(err)
	}

	// http status id OK
	assert.Equal(t, 200, resp.StatusCode)

	// I have four statistics
	assert.Equal(t, 4, len(stats))

	// they are sorted by hits desc
	assert.Equal(t, "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 15]", stats[0].Query)
	assert.Equal(t, "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]", stats[1].Query)
	assert.Equal(t, "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 30]", stats[2].Query)
	assert.Equal(t, "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 30]", stats[3].Query)
}

func Test_it_should_not_display_all_statistics_when_a_limit_is_set(t *testing.T) {
	client := &http.Client{}

	// I save statistics for several queries
	r1, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		resp, err := client.Do(r1)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	r2, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz?word1=hello&word2=world", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 12; i++ {
		resp, err := client.Do(r2)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	r3, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz?limit=30", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 5; i++ {
		resp, err := client.Do(r3)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	r4, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fizzbuzz?word1=hello&word2=world&limit=30", nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 7; i++ {
		resp, err := client.Do(r4)
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	// I request the statistics without a limit
	r, err := http.NewRequest(http.MethodGet, "http://localhost:8080/stats?limit=3", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// I read the response and unmarshal it
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var stats []statistics.StatisticEntry
	err = json.Unmarshal(bytes, &stats)
	if err != nil {
		t.Fatal(err)
	}

	// http status id OK
	assert.Equal(t, 200, resp.StatusCode)

	// I have three statistics
	assert.Equal(t, 3, len(stats))

	// they are sorted by hits desc
	assert.Equal(t, "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 15]", stats[0].Query)
	assert.Equal(t, "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]", stats[1].Query)
	assert.Equal(t, "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 30]", stats[2].Query)
}
