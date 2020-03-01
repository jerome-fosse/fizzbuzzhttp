package statistics

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// StatisticEntry collected for each fizzbuzz request
type StatisticEntry struct {
	Query   string
	Hits    int
	LastHit time.Time
}

func (s StatisticEntry) String() string {
	return fmt.Sprintf("[Hits = %d, Last Hit = %v]", s.Hits, s.LastHit)
}

// StatisticRepository an in memory repository that store statistics in a map. Statistics
// are lost when server stops.
type StatisticRepository struct {
	sync.RWMutex
	stats map[string]StatisticEntry
}

// Save access statistics of a request in a map
func (r *StatisticRepository) Save(query string) {
	logrus.Infof("Saving statistic for request %s", query)
	r.RLock()
	stat, ok := r.stats[query]
	r.RUnlock()

	var newStat StatisticEntry
	if !ok {
		logrus.Debugf("%s is requested for the first time.", query)
		newStat = StatisticEntry{query, 1, time.Now()}
	} else {
		logrus.Debugf("Statistics for request %s were %v", query, stat)
		newStat = StatisticEntry{query, stat.Hits + 1, time.Now()}
	}

	r.Lock()
	r.stats[query] = newStat
	r.Unlock()
}

// FindTopLimitBy returns the x requests the most used.
func (r *StatisticRepository) FindTopLimitBy(limit int) []StatisticEntry {
	values := []StatisticEntry{}
	r.RLock()
	for _, v := range r.stats {
		values = append(values, v)
	}
	r.RUnlock()

	sort.Slice(values, func(i, j int) bool {
		return values[i].Hits > values[j].Hits
	})

	if len(values) < limit {
		limit = len(values)
	}

	retvals := make([]StatisticEntry, limit)
	for i := 0; i < limit; i++ {
		retvals[i] = values[i]
	}

	return retvals
}

// NewRepository create a new StatisticRepository
func NewRepository() *StatisticRepository {
	return &StatisticRepository{stats: make(map[string]StatisticEntry)}
}
