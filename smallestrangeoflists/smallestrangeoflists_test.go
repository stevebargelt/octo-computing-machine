package smallestrangeoflists

import "testing"

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestGetSmallestRange(t *testing.T) {

	list1 := []int{4, 10, 15, 24, 26}
	list2 := []int{0, 9, 12, 20}
	list3 := []int{5, 18, 22, 30}
	lists := [][]int{list1, list2, list3}
	correctResult := []int{20, 24}
	correctResultInt := 5

	result, resultInt, err := GetSmallestRange(lists)
	if err != nil {
		t.Errorf("received error %v", err.Error)
	}
	if !testEq(result, correctResult) {
		t.Error(
			"For", "these lists",
			"expected", correctResult,
			"got", result,
		)
	}
	if resultInt != correctResultInt {
		t.Error(
			"For", "these lists",
			"expected", correctResultInt,
			"got", resultInt,
		)
	}
}

func TestGetSmallestRangeNegative(t *testing.T) {

	list1 := []int{-99, -7, -1, 24, 26}
	list2 := []int{-100, -22, 20}
	list3 := []int{-101, -17, 2, 77, 99}
	lists := [][]int{list1, list2, list3}
	correctResult := []int{-101, -99}
	correctResultInt := 3

	result, resultInt, err := GetSmallestRange(lists)
	if err != nil {
		t.Errorf("received error %v", err.Error)
	}
	if !testEq(result, correctResult) {
		t.Error(
			"For", "these lists",
			"expected", correctResult,
			"got", result,
		)
	}
	if resultInt != correctResultInt {
		t.Error(
			"For", "these lists",
			"expected", correctResultInt,
			"got", resultInt,
		)
	}
}

func TestNilInput(t *testing.T) {

	list1 := []int{4, 10, 15, 24, 26}
	list2 := []int{0, 9, 12, 20}
	lists := [][]int{list1, list2}

	expectedError := "need at least three lists to do this thing"

	_, _, err := GetSmallestRange(lists)
	if err.Error() != expectedError {
		t.Error(
			"For", "these lists",
			"expected", expectedError,
			"got", err.Error(),
		)
	}

}

func TestInputTooSmall(t *testing.T) {

	var lists [][]int
	expectedError := "need at least three lists to do this thing"

	_, _, err := GetSmallestRange(lists)
	if err.Error() != expectedError {
		t.Error(
			"For", "these lists",
			"expected", expectedError,
			"got", err.Error(),
		)
	}

}
