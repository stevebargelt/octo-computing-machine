package finddupinarray

func findDuplicate(array []int) int {

	var actualSum int

	for _, num := range array {
		actualSum += num
	}

	lastItem := array[len(array)-1]

	shouldBe := lastItem * (lastItem + 1) / 2
	return actualSum - shouldBe

}
