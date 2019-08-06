package tests

import (
	"go-google/problems"
	"testing"
)

type rapInput struct {
	n            int
	processes    []string
	allocatedArr []int
	requiredArr  []int
	available    int
	k            int
	ans          int
}

func TestFindMaxAvailableResources(t *testing.T) {
	n := 4
	processes := []string{"p1", "p2", "p3", "p4"}
	allocatedArr := []int{1, 2, 3, 4}
	requiredArr := []int{5, 2, 1, 5}
	available := 1
	k := 2
	input1 := rapInput{n, processes, allocatedArr, requiredArr, available, k, 6}
	testSet := []rapInput{input1}

	for _, test := range testSet {
		if test.ans != problems.FindMaxAvailableResources(test.n, test.processes, test.allocatedArr, test.requiredArr, test.available, test.k) {
			t.Errorf("Test case failure %v", test)
		}
	}

}
