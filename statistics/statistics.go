package statistics

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Statistics collected for each fizzbuzz request
type Statistics struct {
	Hits    int
	LastHit time.Time
}

func (s Statistics) String() string {
	return fmt.Sprintf("[Hits = %d, Last Hit = %v]", s.Hits, s.LastHit)
}

// StatisticRepository an in memory repository that store statistics in a map. Statistics
// are lost when server stops.
type StatisticRepository struct {
	sync.RWMutex
	stats map[string]Statistics
}

// Save access statistics of a request in a map
func (r *StatisticRepository) Save(query string) {
	logrus.Infof("Saving statistic for request %s", query)
	r.RLock()
	stat, ok := r.stats[query]
	r.RUnlock()

	var newStat Statistics
	if !ok {
		logrus.Debugf("%s is requested for the first time.", query)
		newStat = Statistics{1, time.Now()}
	} else {
		logrus.Debugf("Statistics for request %s were %v", query, stat)
		newStat = Statistics{stat.Hits + 1, time.Now()}
	}

	r.Lock()
	r.stats[query] = newStat
	r.Unlock()
}

// FindTopLimitBy returns the x requests the most used.
func (r *StatisticRepository) FindTopLimitBy(limit int) {

}

// NewRepository create a new StatisticRepository
func NewRepository() *StatisticRepository {
	return &StatisticRepository{stats: make(map[string]Statistics)}
}
