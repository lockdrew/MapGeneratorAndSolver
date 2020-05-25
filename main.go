package main

import (
	"fmt"
	"math/rand"
	"time"
)

type maze struct {
	hWalls  [][]bool
	vWalls  [][]bool
	visited [][]bool
	height  int
	width   int
}

type Direction int

const (
	up Direction = iota
	down
	left
	right
)

var Directions = []Direction{
	up,
	down,
	left,
	right,
}

func main() {
	var maze = newMaze(5, 5)
	//maze.print()
	recursiveBackTrackGenerate(maze, 0, 0)

	maze.print()
}

func newMaze(height, width int) *maze {
	var hWalls = make([][]bool, height)
	var vWalls = make([][]bool, height)
	var visited = make([][]bool, height)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			hWalls[h] = append(hWalls[h], true)    // Start with All Horizontal Walls
			vWalls[h] = append(vWalls[h], true)    // Start with All Viritcal Walls
			visited[h] = append(visited[h], false) // Start with All nodes not visited
		}
	}

	return &maze{hWalls, vWalls, visited, height, width}
}

func (maze *maze) print() {
	const hWallLineEnd = "+\n"
	const hWall = "+---"
	const emptyHWall = "+   "
	const vWallLineEnd = "|\n"
	const vWall = "|   "
	const emptyVWall = "    "

	fmt.Print("\n")
	for h := 0; h < maze.height; h++ {
		horizontalWallsLine := ""
		virticalWallsLine := ""
		for w := 0; w < maze.width; w++ {
			horizontalWallsLine += printWall(maze.hWalls[w][h], hWall, emptyHWall)
			virticalWallsLine += printWall(maze.vWalls[w][h], vWall, emptyVWall)
		}
		fmt.Print(horizontalWallsLine + hWallLineEnd)
		fmt.Print(virticalWallsLine + vWallLineEnd)
	}

	finalLine := ""
	for w := 0; w < maze.width; w++ {
		finalLine += hWall
	}
	fmt.Print(finalLine + hWallLineEnd)
	fmt.Print("\n")
}

func printWall(displayWall bool, wall string, emptyWall string) string {
	if displayWall {
		return wall
	}
	return emptyWall
}

func recursiveBackTrackGenerate(maze *maze, x int, y int) {

	fmt.Println("Current Node ", x, y)
	maze.visited[x][y] = true

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Directions), func(i, j int) {
		Directions[i], Directions[j] = Directions[j], Directions[i]
	})

	//Go through all child nodes
	for _, direction := range Directions {
		//time.Sleep(1 * time.Second)

		var newX, newY int
		switch direction {
		case up:
			newX = x
			newY = y - 1
			fmt.Println("up")
			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
				maze.hWalls[newX][newY+1] = false
				fmt.Println("Removing = ", newX, newY)
				recursiveBackTrackGenerate(maze, newX, newY)
			}
		case down:
			newX = x
			newY = y + 1
			fmt.Println("down")
			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
				maze.hWalls[newX][newY] = false
				fmt.Println("Removing = ", newX, newY)
				recursiveBackTrackGenerate(maze, newX, newY)
			}
		case left:
			newX = x - 1
			newY = y
			fmt.Println("left")
			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
				maze.vWalls[newX+1][newY] = false
				fmt.Println("Removing = ", newX, newY)
				recursiveBackTrackGenerate(maze, newX, newY)
			}
		case right:
			newX = x + 1
			newY = y
			fmt.Println("right")
			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
				maze.vWalls[newX][newY] = false
				fmt.Println("Removing = ", newX, newY)
				recursiveBackTrackGenerate(maze, newX, newY)
			}
		default:
			return
		}
	}

	fmt.Println("BackTrack")
}

func (maze *maze) validNode(x, y int) bool {
	validHeight := y < maze.height && y >= 0
	validWidth := x < maze.width && x >= 0
	if validHeight && validWidth {
		return true
	}
	return false
}

func (maze *maze) notvisited(x, y int) bool {
	fmt.Println("Visted - ", maze.visited[x][y])
	return !maze.visited[x][y]
}

func recursiveBackTrackSolve(maze *maze, x, y int) {

}
