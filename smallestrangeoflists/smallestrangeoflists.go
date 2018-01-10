package smallestrangeoflists

import (
	"errors"
	"fmt"
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
// Smallest Range is: [20,22,24] Range of 5 (Last - First +1)
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

	if lists == nil {
		return nil, 0, errors.New("list cannot be nil.")
	}
	if len(lists) < 3 {
		return nil, 0, errors.New("Need at least three lists to do this thing")
	}

	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	var rangeLists []*RangeList
	var minRange []int
	minRangeDiff := MaxInt

	for _, l := range lists {
		rangeLists = append(rangeLists, &RangeList{list: l, currentIndex: 0})
	}

	possibleNext := true

	for possibleNext {
		var workingLists []workinglist
		for i := range rangeLists {
			var wl workinglist
			wl.val = rangeLists[i].Value()
			wl.list = rangeLists[i]
			workingLists = append(workingLists, wl)
		}
		// workingLists[0].list.MoveNext()
		// workingLists[0].list.MoveNext()
		// workingLists[0].list.MoveNext()
		sort.Slice(workingLists, func(i, j int) bool { return workingLists[i].val < workingLists[j].val })
		//workingLists = SortWorkingList(workingLists)
		fmt.Printf("WorkingLists = %v\n", workingLists)
		first := workingLists[0].val
		last := workingLists[len(workingLists)-1].val
		potentialMin := last - first + 1
		if potentialMin < minRangeDiff {
			minRangeDiff = potentialMin
			for i, v := range workingLists {
				if len(minRange) < len(lists) {
					minRange = append(minRange, v.val)
				} else {
					minRange[i] = v.val
				}
			}
		}
		//fmt.Printf("%v", workingLists[0].val)
		possibleNext = workingLists[0].list.MoveNext()
		if !possibleNext {
			fmt.Println("FALSE")
		}
	}
	return minRange, minRangeDiff, nil
}
