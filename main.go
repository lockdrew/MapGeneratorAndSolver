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
	path    [][]bool
	goal    *node
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

type wallDirection int

const (
	virtical wallDirection = iota
	horizontal
)

var directions = []direction{
	up,
	down,
	left,
	right,
}

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	debugColor   = "\033[0;36m%s\033[0m"
)

func main() {
	var maze = newMaze(5, 5)
	recursiveBackTrackGenerate(maze, 0, 0)
	maze.print()

	//fmt.Println("HWall =", maze.hWalls[1][1]) //UpperWall - up is fine... down is not
	//fmt.Println("VWall =", maze.vWalls[1][1]) //leftWAll - left is fine... right is not

	maze.resetVisted()

	maze.goal = &node{3, 3}
	maze.initPath()
	solved := solve(maze, 0, 0)

	maze.printSolution()

	fmt.Println("Solved :", solved)
	// maze.goal = &node{2, 2}
	// fmt.Println("Goal : ", maze.goal.x, maze.goal.y)
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

	return &maze{hWalls, vWalls, visited, height, width, nil, nil}
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

	//fmt.Println("Current Node ", x, y)
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
			//fmt.Println("up")
			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
				maze.hWalls[newX][newY+1] = false
				//fmt.Println("Removing = ", newX, newY)
				recursiveBackTrackGenerate(maze, newX, newY)
			}
		case down:
			newX = x
			newY = y + 1
			//fmt.Println("down")
			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
				maze.hWalls[newX][newY] = false
				//fmt.Println("Removing = ", newX, newY)
				recursiveBackTrackGenerate(maze, newX, newY)
			}
		case left:
			newX = x - 1
			newY = y
			//fmt.Println("left")
			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
				maze.vWalls[newX+1][newY] = false
				//fmt.Println("Removing = ", newX, newY)
				recursiveBackTrackGenerate(maze, newX, newY)
			}
		case right:
			newX = x + 1
			newY = y
			//fmt.Println("right")
			if maze.validNode(newX, newY) && maze.notvisited(newX, newY) {
				maze.vWalls[newX][newY] = false
				//fmt.Println("Removing = ", newX, newY)
				recursiveBackTrackGenerate(maze, newX, newY)
			}
		default:
			return
		}
	}

	//fmt.Println("BackTrack")
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
	//fmt.Println("Visted - ", maze.visited[x][y])
	return !maze.visited[x][y]
}

func (maze *maze) resetVisted() {
	for w := 0; w < maze.width; w++ {
		for h := 0; h < maze.height; h++ {
			maze.visited[w][h] = false
		}
	}
}

func (maze *maze) initPath() {
	maze.path = make([][]bool, maze.height)
	for h := 0; h < maze.height; h++ {
		for w := 0; w < maze.height; w++ {
			maze.path[h] = append(maze.path[h], false)
		}
	}
}

func solve(maze *maze, x int, y int) bool {

	maze.path[x][y] = true
	maze.visited[x][y] = true

	if x == maze.goal.x && y == maze.goal.y {
		return true
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	//maze.printSolution()
	//time.Sleep(1 * time.Second)

	for _, direction := range directions {

		switch direction {
		case up:
			newX := x
			newY := y + 1
			//fmt.Println("current Node", newX, newY)
			if maze.validNode(newX, newY) { //TODO: Apply this change to all if short circuit doesnt work
				//fmt.Println("HWall =", maze.hWalls[newX][newY])
				//fmt.Println("Up")
				if !maze.visited[newX][newY] && !maze.hWalls[newX][newY] { //y+1 is giving problems
					if solve(maze, newX, newY) {
						return true
					}
				}
			}
		case down:
			newX := x
			newY := y - 1
			//fmt.Println("current Node", newX, newY)
			if maze.validNode(newX, newY) {
				//fmt.Println("HWall =", maze.hWalls[newX][newY])
				//fmt.Println("Down")
				if !maze.visited[newX][newY] && !maze.hWalls[newX][newY+1] {

					if solve(maze, newX, newY) {
						return true
					}
				}
			}
		case left:
			newX := x - 1
			newY := y
			//fmt.Println("current Node", newX, newY)
			if maze.validNode(newX, newY) {
				//fmt.Println("vWall =", maze.vWalls[newX][newY])
				//fmt.Println("Left")
				if !maze.visited[newX][newY] && !maze.vWalls[newX+1][newY] {
					if solve(maze, newX, newY) {
						return true
					}
				}
			}
		case right:
			newX := x + 1
			newY := y
			//fmt.Println("current Node", newX, newY)
			if maze.validNode(newX, newY) {
				//fmt.Println("vWall =", maze.vWalls[newX][newY])
				//fmt.Println("right")
				if !maze.visited[newX][newY] && !maze.vWalls[newX][newY] {
					if solve(maze, newX, newY) {
						return true
					}
				}
			}

		default:
			//fmt.Println("default")
			return false
		}
	}

	//fmt.Println("BackTrack")

	maze.path[x][y] = false
	return false
}

func (maze *maze) printSolution() {
	const hWallLineEnd = "+\n"
	const hWall = "+---"
	const emptyHWall = "+   "
	const vWallLineEnd = "|\n"
	const vWall = "|   "
	const vWallPath = "| * " //Color text output
	const emptyVWall = "    "
	const emptyVWallPath = " * " //Color text output
	// const goal = "  x  "
	// const goalWithWall = "|  x  "

	fmt.Print("\n")
	for h := 0; h < maze.height; h++ {
		horizontalWallsLine := ""
		virticalWallsLine := ""
		for w := 0; w < maze.width; w++ {
			horizontalWallsLine += maze.printWallSol(horizontal, w, h)
			virticalWallsLine += maze.printWallSol(virtical, w, h)
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

//TODO: Add logic w/ statement to determine  which to print
func (maze *maze) printHorizontalWall(x, y int) string {
	const hWall = "+---"
	const emptyHWall = "+   "

	if maze.hWalls[x][y] {
		return hWall
	}
	return emptyHWall
}

func (maze *maze) printViriticalWall(x, y int) string {

	const vWall = "|   "
	const vWallPath = "| * "
	const emptyVWall = "    "
	const emptyVWallPath = "  * "
	const emptyGoal = "  X "
	const GoalVWall = "| X "

	if maze.vWalls[x][y] {

		if maze.goal.x == x && maze.goal.y == y {
			return GoalVWall
		}

		if maze.path[x][y] {
			return vWallPath
		}
		return vWall
	}

	if maze.path[x][y] {

		if maze.goal.x == x && maze.goal.y == y {
			return emptyGoal
		}

		return emptyVWallPath
	}

	return emptyVWall
}

func (maze *maze) printWallSol(direction wallDirection, x int, y int) string {
	if direction == virtical {
		return maze.printViriticalWall(x, y)
	}
	return maze.printHorizontalWall(x, y)
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
