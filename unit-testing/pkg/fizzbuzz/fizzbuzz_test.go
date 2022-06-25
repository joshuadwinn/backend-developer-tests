package fizzbuzz

import (
	"testing"
	"reflect"
)
func TestFizzBuzz(t *testing.T) {

	var values []string
	values = FizzBuzz(20, 3, 5)
	results := []string {"1", "2", "Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz","16"}
	reflect.DeepEqual(values, results)

	defer func() { recover() }()
	FizzBuzz(-20, 3, 5)
	t.Errorf("Should have panicked")
}
