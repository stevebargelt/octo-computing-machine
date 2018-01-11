package islandcount

import "errors"

// Given a 2D map of 1s (Land) and 0s (Water), count the nuber of islands.
// An island is surrounded by water. Assume grid is surrounded by water.
//
//
// 11110
// 11010
// 11000
// 00000
//
// Answer: 1

// At least 2 rows
// Each row has equal columns
// JUST 1s and 0s
//

// GetIslandCount returns the number of islands in a grid
func GetIslandCount(grid [][]byte) (int, error) {
	if len(grid) < 2 {
		return 0, errors.New("must have more than two rows")
	}
	islandCount := 0
	var lastRow []byte
	for _, row := range grid {
		var lastCoord byte
		for i, coord := range row {
			thisIsLand := coord == 1
			continuingIsland := false
			if thisIsLand && lastCoord == 1 {
				continuingIsland = true
			}
			if lastRow != nil && lastRow[i] == 1 {
				continuingIsland = true
			}

			if thisIsLand && !continuingIsland {
				islandCount++
			}
			lastCoord = coord

		}
		lastRow = row
	}

	return islandCount, nil
}
