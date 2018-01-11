package islandcount

import "testing"

func TestIslandCount(t *testing.T) {

	row1 := []byte{1, 1, 1, 1, 0}
	row2 := []byte{1, 1, 0, 1, 0}
	row3 := []byte{1, 1, 0, 0, 0}
	row4 := []byte{0, 0, 0, 0, 0}
	grid := [][]byte{row1, row2, row3, row4}
	expected := 1
	actual, _ := GetIslandCount(grid)

	if expected != actual {
		t.Error(
			"For", "these lists",
			"expected", expected,
			"got", actual,
		)
	}

}

func TestIslandCountQbikez(t *testing.T) {

	row1 := []byte{1, 1, 0, 0, 0}
	row2 := []byte{1, 1, 0, 0, 0}
	row3 := []byte{0, 0, 1, 0, 1}
	row4 := []byte{0, 0, 1, 1, 1}
	grid := [][]byte{row1, row2, row3, row4}
	expected := 2
	actual, _ := GetIslandCount(grid)

	if expected != actual {
		t.Error(
			"For", "these lists",
			"expected", expected,
			"got", actual,
		)
	}

}
