package main

import "testing"

func TestServerConnection(t *testing.T) {

	client := connectToServer()
	if client == nil {
		t.Log("Failed to connect")
		t.Fail()
	}
}

func TestAddNumbers(t *testing.T) {
	client := connectToServer()
	result := addNumbers(1, 2, client)
	if result != 3 {
		t.Log("Failed to add")
		t.Fail()
	}
}

func TestSubtractNumbers(t *testing.T) {
	client := connectToServer()
	result := subtractNumbers(5, 2, client)
	if result != 3 {
		t.Log("Failed to add")
		t.Fail()
	}
}

func TestCalculateAvrage(t *testing.T) {
	client := connectToServer()
	numbers := []float32{1, 5, 6}
	result := calculateAverage(numbers, client)
	if result != 4 {
		t.Log("Failed to add")
		t.Fail()
	}
}
