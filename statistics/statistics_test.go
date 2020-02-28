package statistics

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_query_is_saved(t *testing.T) {
	// Given a FizzBuzzQueryParameters
	params := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]"

	// When I save the statistics for the first time
	repo := NewRepository()
	repo.Save(params)

	// Then statistics entry is created
	val, ok := repo.stats[params]
	assert.True(t, ok)

	// And hits = 1 and LastHit is today
	assert.Equal(t, 1, val.Hits)
	y1, m1, d1 := time.Now().Date()
	y2, m2, d2 := val.LastHit.Date()
	assert.Equal(t, y1, y2)
	assert.Equal(t, m1, m2)
	assert.Equal(t, d1, d2)
}

func Test_statistics_should_be_updated_each_time_the_same_query_is_saved_several_times(t *testing.T) {
	// Given a FizzBuzzQueryParameters
	params := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]"

	// When I save the statistics for several times
	repo := NewRepository()
	repo.Save(params)
	firstVal := repo.stats[params]

	n := rand.Intn(8) + 1
	for i := 0; i < n; i++ {
		repo.Save(params)
		time.Sleep(100 * time.Millisecond)
	}

	// Then statistics entry is created
	val, ok := repo.stats[params]
	assert.True(t, ok)

	// And hits is equal to the number of time statistics have been saved
	assert.Equal(t, n+1, val.Hits)

	// And LastHit is after FirstHit
	assert.True(t, val.LastHit.After(firstVal.LastHit))
}
