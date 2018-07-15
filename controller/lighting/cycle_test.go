package lighting

import (
	"testing"
	"time"
)

func TestGetValue(t *testing.T) {
	intensities := []int{0, 0, 0, 0, 20, 40, 60, 70, 30, 10, 0, 0}
	ch := Channel{
		Values: intensities,
	}
	if err := ValidateValues(intensities); err != nil {
		t.Fatal(err)
	}
	t1 := time.Date(2015, time.October, 12, 7, 0, 0, 0, time.UTC)
	v := ch.GetValue(t1)
	if v != 10 {
		t.Fatal("Expected 10, found:", v)
	}
}
