package fizzbuzzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Values_multiples_of_int1_should_be_replaced_by_word1(t *testing.T) {
	// Given a fizzbuzzer that replace all values divisible by three only by fizz
	fb := New(3, 5, 6, "fizz", "buzz").(*fizzbuzzer)

	// All values divisibles by 3 only are replaced by fizz
	assert.Equal(t, "fizz", fb.getFizzbuzzValue(3))
	assert.Equal(t, "fizz", fb.getFizzbuzzValue(6))
	assert.Equal(t, "fizz", fb.getFizzbuzzValue(9))
	assert.Equal(t, "fizz", fb.getFizzbuzzValue(12))
	assert.NotEqual(t, "fizz", fb.getFizzbuzzValue(15))
}

func Test_Values_multiples_of_int2_should_be_replaced_by_word2(t *testing.T) {
	// Given a fizzbuzzer that replace all values divisible by five only by buzz
	fb := New(3, 5, 6, "fizz", "buzz").(*fizzbuzzer)

	// All values divisibles by 5 only are replaced y buzzb
	assert.Equal(t, "buzz", fb.getFizzbuzzValue(5))
	assert.Equal(t, "buzz", fb.getFizzbuzzValue(10))
	assert.NotEqual(t, "buzz", fb.getFizzbuzzValue(15))
	assert.Equal(t, "buzz", fb.getFizzbuzzValue(20))
	assert.Equal(t, "buzz", fb.getFizzbuzzValue(25))
}

func Test_values_divisible_by_int1_and_int2_are_replaced_by_word1word2(t *testing.T) {
	// Given a fizzbuzzer that replace all values divisible by three only by fizz and all values divisible by five only by buzz
	fb := New(3, 5, 6, "fizz", "buzz").(*fizzbuzzer)

	// All values divisibles by 3 and 5 are replaced by fizzbuzz
	assert.Equal(t, "fizzbuzz", fb.getFizzbuzzValue(15))
	assert.Equal(t, "fizzbuzz", fb.getFizzbuzzValue(30))
	assert.Equal(t, "fizzbuzz", fb.getFizzbuzzValue(45))
}

func Test_Values_divisible_nor_by_int1_and_int2_are_converted_to_a_string(t *testing.T) {
	// Given a fizzbuzzer that replace all values divisible by three only by fizz and all values divisible by five only by buzz
	fb := New(3, 5, 6, "fizz", "buzz").(*fizzbuzzer)

	// All values divisibles by 5 only are replaced by buzz
	assert.Equal(t, "1", fb.getFizzbuzzValue(1))
	assert.Equal(t, "2", fb.getFizzbuzzValue(2))
	assert.NotEqual(t, "3", fb.getFizzbuzzValue(3))
	assert.Equal(t, "4", fb.getFizzbuzzValue(4))
	assert.NotEqual(t, "5", fb.getFizzbuzzValue(5))
	assert.NotEqual(t, "15", fb.getFizzbuzzValue(15))
}

func Test_Fizzbuzzer_should_return_as_many_strings_as_the_limit(t *testing.T) {
	// Given a default fizzbuzzer
	fb := NewWithDefaultValues().(*fizzbuzzer)

	// When I get the fizzbuzz values
	values := fb.Get()

	// Then I the number of values is equal to the limit of the fizzbuzzer
	assert.Equal(t, fb.limit, len(values))
}

func Test_Fizzbuzzer_should_generate_a_correct_fizzbuzz(t *testing.T) {
	// Given a fizzbuzzer with a limit of 15
	fb := New(3, 5, 15, "fizz", "buzz")

	// When I get the fizzbuzz values
	values := fb.Get()

	// Then values should be equal to [1 2 fizz 4 buzz fizz è _ fizz buzz 11 fizz 13 14 fizzbuzz]
	assert.Equal(t, []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}, values)
}
