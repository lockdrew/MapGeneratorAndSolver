package main

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	x int
	y int
}

type maze struct {
	hWalls  [][]bool
	vWalls  [][]bool
	visited [][]bool
	height  int
	width   int
	goal    *node
}

type solution struct {
	visited [][]bool
	path    [][]bool
	goalX   int
	goalY   int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

var directions = []direction{
	up,
	down,
	left,
	right,
}

func main() {
	var maze = newMaze(5, 5)
	recursiveBackTrackGenerate(maze, 0, 0)
	maze.print()

	maze.goal = &node{2, 2}
	fmt.Println("Goal : ", maze.goal.x, maze.goal.y)
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

	return &maze{hWalls, vWalls, visited, height, width, nil}
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
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	//Go through all child nodes
	for _, direction := range directions {
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

func (maze *maze) resetVisted() {
	for w := 0; w < maze.width; w++ {
		for h := 0; h < maze.height; h++ {
			maze.visited[w][h] = false
		}
	}
}

func solve(solution *solution, maze *maze, x int, y int) bool {
	if x == solution.goalX && y == solution.goalY {
		return true
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	for _, direction := range directions {
		switch direction {
		case up:
			newX := x
			newY := y + 1
			if solution.visited[newX][newY] == false && maze.validNode(newX, newY) && maze.hWalls[newX][newY+1] == false {
				solution.visited[newX][newY] = true
				solution.path[newX][newY] = true
				if solve(solution, maze, newX, newY) {
					return true
				}
			}
		case down:
			newX := x
			newY := y + 1
			if solution.visited[newX][newY] == false && maze.validNode(newX, newY) && maze.hWalls[newX][newY] == false {
				solution.visited[newX][newY] = true
				solution.path[newX][newY] = true
				if solve(solution, maze, newX, newY) {
					return true
				}
			}
		case left:
			newX := x - 1
			newY := y
			if solution.visited[newX][newY] == false && maze.validNode(newX, newY) && maze.vWalls[newX+1][newY] == false {
				solution.visited[newX][newY] = true
				solution.path[newX][newY] = true
				if solve(solution, maze, newX, newY) {
					return true
				}
			}
		case right:
		default:
			return false
		}
	}

	solution.path[x][y] = false
	return false
}

// func recursiveBackTrackSolve(maze *maze, x, y int) {
// 	if maze.goal.x == x && maze.goal.y == y {
// 		return
// 	}

// 	//Go through all child nodes
// 	for _, direction := range directions {
// 		//time.Sleep(1 * time.Second)

// 		var newX, newY int
// 		switch direction {
// 		case up:
// 			newX = x
// 			newY = y - 1
// 			fmt.Println("up")
// 			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
// 				maze.hWalls[newX][newY+1] = false
// 				fmt.Println("Removing = ", newX, newY)
// 				recursiveBackTrackGenerate(maze, newX, newY)
// 			}
// 		case down:
// 			newX = x
// 			newY = y + 1
// 			fmt.Println("down")
// 			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
// 				maze.hWalls[newX][newY] = false
// 				fmt.Println("Removing = ", newX, newY)
// 				recursiveBackTrackGenerate(maze, newX, newY)
// 			}
// 		case left:
// 			newX = x - 1
// 			newY = y
// 			fmt.Println("left")
// 			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
// 				maze.vWalls[newX+1][newY] = false
// 				fmt.Println("Removing = ", newX, newY)
// 				recursiveBackTrackGenerate(maze, newX, newY)
// 			}
// 		case right:
// 			newX = x + 1
// 			newY = y
// 			fmt.Println("right")
// 			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
// 				maze.vWalls[newX][newY] = false
// 				fmt.Println("Removing = ", newX, newY)
// 				recursiveBackTrackGenerate(maze, newX, newY)
// 			}
// 		default:
// 			return
// 		}
// 	}
// }
