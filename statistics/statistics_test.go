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

func Test_it_should_return_statistics_sorted_by_hits(t *testing.T) {
	// Given 4 queries saved with their statistics
	query1 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]"
	query2 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 15]"
	query3 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 30]"
	query4 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 30]"

	repo := NewRepository()
	repo.stats[query1] = StatisticEntry{query1, 10, time.Now()}
	repo.stats[query2] = StatisticEntry{query2, 5, time.Now()}
	repo.stats[query3] = StatisticEntry{query3, 12, time.Now()}
	repo.stats[query4] = StatisticEntry{query4, 7, time.Now()}

	// When I get the most asked queries
	top := repo.FindTopLimitBy(4)

	// I have 4 elements in my top
	assert.Equal(t, 4, len(top))

	// And they are sorted by hits desc
	assert.Equal(t, query3, top[0].Query)
	assert.Equal(t, query1, top[1].Query)
	assert.Equal(t, query4, top[2].Query)
	assert.Equal(t, query2, top[3].Query)
}

func Test_it_should_return_first_three_statistics_out_of_four(t *testing.T) {
	// Given 4 queries saved with their statistics
	query1 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]"
	query2 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 15]"
	query3 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 30]"
	query4 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 30]"

	repo := NewRepository()
	repo.stats[query1] = StatisticEntry{query1, 10, time.Now()}
	repo.stats[query2] = StatisticEntry{query2, 5, time.Now()}
	repo.stats[query3] = StatisticEntry{query3, 12, time.Now()}
	repo.stats[query4] = StatisticEntry{query4, 7, time.Now()}

	// When I get the most asked queries
	top := repo.FindTopLimitBy(3)

	// I have 3 elements in my top
	assert.Equal(t, 3, len(top))

	// And they are sorted by hits desc
	assert.Equal(t, query3, top[0].Query)
	assert.Equal(t, query1, top[1].Query)
	assert.Equal(t, query4, top[2].Query)
}

func Test_it_should_return_all_statistics_when_I_request_more_than_the_number_of_entries(t *testing.T) {
	// Given 4 queries saved with their statistics
	query1 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 15]"
	query2 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 15]"
	query3 := "[int1 = 3, word1 = fizz, int2 = 5, word2 = buzz, limit = 30]"
	query4 := "[int1 = 3, word1 = hello, int2 = 5, word2 = world, limit = 30]"

	repo := NewRepository()
	repo.stats[query1] = StatisticEntry{query1, 10, time.Now()}
	repo.stats[query2] = StatisticEntry{query2, 5, time.Now()}
	repo.stats[query3] = StatisticEntry{query3, 12, time.Now()}
	repo.stats[query4] = StatisticEntry{query4, 7, time.Now()}

	// When I get the most asked queries
	top := repo.FindTopLimitBy(10)

	// I have 3 elements in my top
	assert.Equal(t, 4, len(top))

	// And they are sorted by hits desc
	assert.Equal(t, query3, top[0].Query)
	assert.Equal(t, query1, top[1].Query)
	assert.Equal(t, query4, top[2].Query)
	assert.Equal(t, query2, top[3].Query)
}
