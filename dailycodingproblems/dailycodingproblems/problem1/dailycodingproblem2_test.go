package main

import "testing"

func TestCodingProblem(t *testing.T){

	//testing for a valid argument
	result := add(1,2)
	if result != 3{
		t.Errorf("Error addition not working")
	}
}