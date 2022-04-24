package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("A new connection object must have been made! ")
	}

	expectedCounter := counter1

	currentCounter := counter1.AddOne()
	if currentCounter != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCounter)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("Singleton instances must be different")
	}

	currentCounter = counter2.AddOne()
	if currentCounter != 2 {
		t.Errorf("After calling 'AddOne' using  the second counter, the curent count must be 2 but was %d\n", currentCounter)
	}
}
