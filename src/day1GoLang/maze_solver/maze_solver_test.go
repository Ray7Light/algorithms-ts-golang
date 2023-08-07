package mazeSolver

import (
	"reflect"
	"strings"
	"testing"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}

	start := Point{x: 10, y: 0}
	end := Point{x: 1, y: 5}

	result := MazeSolver(maze, "x", start, end)
	expected := drawPath(maze, mazeResult)

	if reflect.DeepEqual(result, expected) == false {
		t.Errorf("the solution is wrong")
	}
}

func drawPath(data []string, path []Point) []string {
	var data2 [][]string

	for _, row := range data {
		data2 = append(data2, strings.Split(row, ""))
	}

	for _, p := range path {
		if p.y < len(data2) && p.x < len(data2[p.y]) {
			data2[p.y][p.x] = "*"
		}
	}

	var result []string

	for _, row := range data2 {
		result = append(result, strings.Join(row, ""))
	}

	return result
}
