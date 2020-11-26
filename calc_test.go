package calc

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
		//t.Errorf("Error in values")
	}

	result = Add(1, 0)
	if result != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
		//t.Errorf("Error in values")
	}

	result = Add(2, -2)
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
		//t.Errorf("Error in values")
	}
}
