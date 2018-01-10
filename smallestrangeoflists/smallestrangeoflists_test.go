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
	correctResult := []int{20, 22, 24}

	result, err := GetSmallestRange(lists)
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
}
