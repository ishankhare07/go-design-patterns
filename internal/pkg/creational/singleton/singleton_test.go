package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter := GetInstance()

	if counter == nil {
		t.Errorf("Expected pointer to Singleton. Got nil.")
	}

	currentCount := counter.AddOne()
	if currentCount != 1 {
		t.Errorf("Counter value must be 1 after first call. Got = %d instead.", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != counter {
		t.Errorf("Expected same instance in counter2, got a different instance instead.")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("Counter after second call must be 2. Got %d instead", currentCount)
	}
}
