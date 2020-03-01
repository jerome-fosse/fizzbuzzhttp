package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/object-it/fizzbuzzhttp/statistics"
	"github.com/stretchr/testify/assert"
)

func Test_when_a_limit_is_set_it_should_return_the_number_of_statistics_corresponding(t *testing.T) {
	// Given existing statistics for four different queries
	repo := statistics.NewRepository()
	query1 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]"
	query2 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 15]"
	query3 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 30]"
	query4 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 30]"

	for i := 0; i < 10; i++ {
		repo.Save(query1)
	}
	for i := 0; i < 5; i++ {
		repo.Save(query2)
	}
	for i := 0; i < 12; i++ {
		repo.Save(query3)
	}
	for i := 0; i < 7; i++ {
		repo.Save(query4)
	}

	// When I request GET /stats?limit=4
	req, err := http.NewRequest(http.MethodGet, "/stats?limit=3", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(NewStatsHandler(repo))
	handler.ServeHTTP(rec, req)

	// Then the http status code is OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// And I get the number of statistics corresponding to the limit
	var stats []statistics.StatisticEntry
	err = json.Unmarshal(rec.Body.Bytes(), &stats)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(stats))
	assert.Equal(t, query3, stats[0].Query)
	assert.Equal(t, query1, stats[1].Query)
	assert.Equal(t, query4, stats[2].Query)
}

func Test_when_a_limit_is_not_set_it_should_return_all_the_statistics(t *testing.T) {
	// Given existing statistics for four different queries
	repo := statistics.NewRepository()
	query1 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]"
	query2 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 15]"
	query3 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 30]"
	query4 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 30]"

	for i := 0; i < 10; i++ {
		repo.Save(query1)
	}
	for i := 0; i < 5; i++ {
		repo.Save(query2)
	}
	for i := 0; i < 12; i++ {
		repo.Save(query3)
	}
	for i := 0; i < 7; i++ {
		repo.Save(query4)
	}

	// When I request GET /stats?limit=4
	req, err := http.NewRequest(http.MethodGet, "/stats?", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(NewStatsHandler(repo))
	handler.ServeHTTP(rec, req)

	// Then the http status code is OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// And I get all the statistics
	var stats []statistics.StatisticEntry
	err = json.Unmarshal(rec.Body.Bytes(), &stats)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 4, len(stats))
	assert.Equal(t, query3, stats[0].Query)
	assert.Equal(t, query1, stats[1].Query)
	assert.Equal(t, query4, stats[2].Query)
	assert.Equal(t, query2, stats[3].Query)
}
