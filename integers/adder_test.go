package integers

import (
	"fmt"
	"testing"
)

// More testing : https://blog.gopheracademy.com/advent-2017/property-based-testing/
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Add function example
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
