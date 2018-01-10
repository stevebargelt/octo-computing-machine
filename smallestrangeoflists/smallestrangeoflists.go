package smallestrangeoflists

import (
	"errors"
	"sort"
)

// You have K lists of sorted integers.
// Find the smalest range that includes at least one number from each of the k lists
//
//
// List 1: [4,10,15,24,26]
// List 2: [0,9,12,20]
// List 3: [5,18,22,30]
//
// Smallest Range is: [20,,24] Range of 5 (Last - First +1)
//

type RangeList struct {
	list         []int
	currentIndex int
}

func (rl *RangeList) Value() int {
	return rl.list[rl.currentIndex]
}

func (rl *RangeList) MoveNext() bool {
	if rl.currentIndex == len(rl.list)-1 {
		return false
	} else {
		rl.currentIndex++
		return true
	}
}

type workinglist struct {
	val  int
	list *RangeList
}

// GetSmallestRange returns the smallest range possible from three or more lists
func GetSmallestRange(lists [][]int) ([]int, int, error) {

	// NOTE : The len and cap functions will both return 0 for a nil slice.
	if lists == nil {
		return nil, 0, errors.New("list cannot be nil")
	}
	if len(lists) < 3 {
		return nil, 0, errors.New("Need at least three lists to do this thing")
	}

	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	var rangeLists []*RangeList
	var minRange [2]int
	minRangeDiff := MaxInt

	for _, l := range lists {
		rangeLists = append(rangeLists, &RangeList{list: l, currentIndex: 0})
	}

	possibleNext := true

	for possibleNext {
		var workingLists []workinglist // yes clear out every time... refilled with the rangeLists
		for i := range rangeLists {    // can't use i, r := range rangeLists since range always copies the object
			workingLists = append(workingLists, workinglist{rangeLists[i].Value(), rangeLists[i]})
		}
		sort.Slice(workingLists, func(i, j int) bool { return workingLists[i].val < workingLists[j].val })
		first := workingLists[0].val
		last := workingLists[len(workingLists)-1].val
		potentialMin := last - first + 1
		if potentialMin < minRangeDiff {
			minRangeDiff = potentialMin
			minRange[0] = first
			minRange[1] = last
		}
		possibleNext = workingLists[0].list.MoveNext()
	}
	return minRange[0:], minRangeDiff, nil
}
