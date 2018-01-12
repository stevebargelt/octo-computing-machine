package islandcount

import (
	"errors"

	"github.com/stevebargelt/missionInterview/stack"
)

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

type point struct {
	X int
	Y int
}

func mapLand(ocean [][]byte, startCoord point) {

	toVisit := stack.New()
	toVisit.Push(startCoord)
	for toVisit.Len() > 0 {
		cur := toVisit.Pop().(point)
		//check all four directions
		//left
		if cur.X > 0 && ocean[cur.X-1][cur.Y] == 1 {
			toVisit.Push(point{X: cur.X - 1, Y: cur.Y})
		}
		//right
		if cur.X < len(ocean)-1 && ocean[cur.X+1][cur.Y] == 1 {
			toVisit.Push(point{X: cur.X + 1, Y: cur.Y})
		}
		//up
		if cur.Y > 0 && ocean[cur.X][cur.Y-1] == 1 {
			toVisit.Push(point{X: cur.X, Y: cur.Y - 1})
		}
		//down
		if cur.Y < len(ocean[cur.X])-1 && ocean[cur.X][cur.Y+1] == 1 {
			toVisit.Push(point{X: cur.X, Y: cur.Y + 1})
		}
		ocean[cur.X][cur.Y] = 2

	}

}

func getUnvisitedLand(ocean [][]byte, lastVisited point) (point, error) {

	for x := lastVisited.X; x < len(ocean); x++ {
		for y := 0; y < len(ocean[x]); y++ {
			if ocean[x][y] == 1 {
				return point{X: x, Y: y}, nil
			}
		}
	}
	return point{X: 0, Y: 0}, errors.New("Visited")
}

// GetIslandCount returns the number of islands in a grid
func GetIslandCount(ocean [][]byte) (int, error) {
	if len(ocean) < 2 {
		return 0, errors.New("must have more than one row")
	}
	start := point{X: 0, Y: 0}
	islandCount := 0
	for {
		start, err := getUnvisitedLand(ocean, start)
		if err != nil {
			break
		}
		mapLand(ocean, start)
		islandCount++

	}
	return islandCount, nil
}
