package mazeSolver

import (
	"strings"
)

type Point struct {
	y int
	x int
}

// Directions: top, right, bottom, left.
var dirs = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func MazeSolver(maze []string, wall string, start, end Point) []Point {
	mazeLen := len(maze)
	seen := make([][]string, mazeLen)
	splitMaze := [][]string{}
	path := []Point{}

	for _, row := range maze {
		splitMaze = append(splitMaze, strings.Split(row, ""))
	}

	for i := range seen {
		seen[i] = make([]string, len(splitMaze[0]))
	}

	for i := 0; i < mazeLen; i++ {
		for j := 0; j < len(splitMaze[0]); j++ {
			seen[i][j] = ""
		}
	}

	walk(splitMaze, wall, seen, start, end, &path)

	return path
}

func walk(maze [][]string, wall string, seen [][]string, curr, end Point, path *[]Point) bool {

	// 1. End condition.
	// Reached the end.
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, curr)
		return true
	}

	// Wall hit or out of bounds.
	if curr.y < 0 || curr.y >= len(maze) || curr.x < 0 || curr.x >= len(maze[0]) || maze[curr.y][curr.x] == wall {
		return false
	}

	// Node already visited.
	if seen[curr.y][curr.x] == "*" {
		return false
	}

	// 2. Recurse
	seen[curr.y][curr.x] = "*"
	*path = append(*path, curr)

	for _, dir := range dirs {
		next := Point{y: curr.y + dir[0], x: curr.x + dir[1]}

		if walk(maze, wall, seen, next, end, path) {
			return true
		}
	}

	// post
	// remove path that did not lead to an exit.
	*path = (*path)[:len(*path)-1]

	return false
}
