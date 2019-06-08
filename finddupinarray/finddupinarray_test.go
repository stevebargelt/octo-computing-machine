package finddupinarray

import (
	"testing"
)

func Test(t *testing.T) {

	testArray := []int{1, 2, 3, 4, 4, 5}
	result := findDuplicate(testArray)

	if result != 4 {
		t.Errorf("result should have been 4")
	}

	testArray = []int{1, 2, 3, 4, 5, 5}
	result = findDuplicate(testArray)

	if result != 5 {
		t.Errorf("result should have been 5")
	}

	testArray = []int{1, 1, 2, 3, 4, 5}
	result = findDuplicate(testArray)

	if result != 1 {
		t.Errorf("result should have been 1")
	}

}
