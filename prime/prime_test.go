package prime

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	primes := GetPrime(10)
	expected := []int{2, 3, 5, 7}

	// Use reflect.DeepEqual to compare slices
	if !reflect.DeepEqual(primes, expected) {
		t.Errorf("expected %v but got %v", expected, primes)
	}
}
