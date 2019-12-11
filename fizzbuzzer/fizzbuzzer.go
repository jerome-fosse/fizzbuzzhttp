package fizzbuzzer

// Fizzbuzzer interface for Fizzbuzz implementations
type Fizzbuzzer interface {
	Get() string
}

type defaultFizzbuzzer struct {}

func (fb *defaultFizzbuzzer) Get() string {
	return "fizzbuzz"
}

// New create a new Fizzbuzzer
func New() Fizzbuzzer {
	return &defaultFizzbuzzer{}
}
