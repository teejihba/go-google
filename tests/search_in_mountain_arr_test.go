package tests

import (
	"go-google/problems"
	"testing"
)

func TestFindIfExists(t *testing.T) {
	arr := []int{1, 5, 9, 20, 24, 36, 48, 50, 12, 8, 6, 4, 3, 2, 1, 0, -5, -8, -15, -19, -23, -26, -29, -31, -35}

	input := problems.Input{arr, len(arr), 0}
	ans1 := input.FindIfExists(-19)
	if ans1 != true {
		t.Errorf("Test case failure for search in mountain arr")
	}

	ans2 := input.FindIfExists(13)
	if ans2 != false {
		t.Errorf("Test case failure for search in mountain arr")
	}

	ans := input.FindIfExists(50)
	if ans != true {
		t.Errorf("Test cases failure for search in mountain arr")
	}

}
