package fizzbuzzer

import (
	"strconv"
)

// Fizzbuzzer interface for Fizzbuzz implementations
type Fizzbuzzer interface {
	Get() []string
}

type fizzbuzzer struct {
	int1, int2, limit int
	word1, word2      string
}

func (fb *fizzbuzzer) Get() []string {
	result := make([]string, fb.limit)

	for i := 0; i < fb.limit; i++ {
		result[i] = fb.getFizzbuzzValue(i + 1)
	}

	return result
}

func (fb *fizzbuzzer) getFizzbuzzValue(i int) string {
	mod1 := i % fb.int1
	mod2 := i % fb.int2

	switch {
	case mod1 == 0 && mod2 == 0:
		return fb.word1 + fb.word2
	case mod1 == 0 && mod2 != 0:
		return fb.word1
	case mod1 != 0 && mod2 == 0:
		return fb.word2
	default:
		return strconv.Itoa(i)
	}
}

// NewWithDefaultValues create a new Fizzbuzzer
func NewWithDefaultValues() Fizzbuzzer {
	return New(3, 5, 15, "Fizz", "Buzz")
}

// New create a fizzbuzzer with the parameters given to the constructor
func New(int1, int2, limit int, word1, word2 string) Fizzbuzzer {
	return &fizzbuzzer{int1: int1, int2: int2, limit: limit, word1: word1, word2: word2}
}
