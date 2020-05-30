package main

import (
	"fmt"
	"math/rand"
	"time"
)

type maze struct {
	hWalls          [][]bool
	vWalls          [][]bool
	generateVisited [][]bool
	solveVisited    [][]bool
	height          int
	width           int
	path            [][]bool
	goalX           int
	goalY           int
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

func main() {
	var maze = newMaze(15, 15, 2, 5)
	maze.generate(0, 0)

	maze.printSolution()

	solved := solve(maze, 0, 0)

	maze.printSolution()

	fmt.Println("Solved :", solved)
}

func newMaze(height, width, goalX, goalY int) *maze {
	var hWalls = make([][]bool, height)
	var vWalls = make([][]bool, height)
	var solveVisited = make([][]bool, height)
	var generateVisited = make([][]bool, height)
	var path = make([][]bool, height)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			hWalls[h] = append(hWalls[h], true)                    // Start with all horizontal walls
			vWalls[h] = append(vWalls[h], true)                    // Start with all viritcal walls
			solveVisited[h] = append(solveVisited[h], false)       // Start with all nodes not visited
			generateVisited[h] = append(generateVisited[h], false) // Start with all nodes not visited
			path[h] = append(path[h], false)                       // Start with empty path
		}
	}

	return &maze{hWalls, vWalls, generateVisited, solveVisited, height, width, path, goalX, goalY}
}

func shuffleDirections() []direction {

	directions := []direction{
		up,
		down,
		left,
		right,
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(directions), func(i, j int) {
		directions[i], directions[j] = directions[j], directions[i]
	})

	return directions
}

func (maze *maze) generate(x int, y int) {

	maze.generateVisited[x][y] = true

	directions := shuffleDirections()

	//Go through all child nodes
	for _, direction := range directions {
		var newX, newY int
		switch direction {
		case up:
			newX = x
			newY = y - 1
			if maze.validNode(newX, newY) && !maze.generateVisited[newX][newY] {
				maze.clearHorzontalWall(newX, newY+1)
				maze.generate(newX, newY)
			}
		case down:
			newX = x
			newY = y + 1
			if maze.validNode(newX, newY) && !maze.generateVisited[newX][newY] {
				maze.clearHorzontalWall(newX, newY)
				maze.generate(newX, newY)
			}
		case left:
			newX = x - 1
			newY = y
			if maze.validNode(newX, newY) && !maze.generateVisited[newX][newY] {
				maze.clearViriticalWall(newX+1, newY)
				maze.generate(newX, newY)
			}
		case right:
			newX = x + 1
			newY = y
			if maze.validNode(newX, newY) && !maze.generateVisited[newX][newY] {
				maze.clearViriticalWall(newX, newY)
				maze.generate(newX, newY)
			}
		default:
			return
		}
	}
}

func (maze *maze) clearViriticalWall(x, y int) {
	maze.vWalls[x][y] = false
}

func (maze *maze) virticalWall(x, y int) bool {
	return maze.vWalls[x][y]
}

func (maze *maze) clearHorzontalWall(x, y int) {
	maze.hWalls[x][y] = false
}

func (maze *maze) horizontalWall(x, y int) bool {
	return maze.hWalls[x][y]
}

func (maze *maze) setPath(x, y int) {
	maze.path[x][y] = true
}

func (maze *maze) unsetPath(x, y int) {
	maze.path[x][y] = false
}

func (maze *maze) validNode(x, y int) bool {
	validHeight := y < maze.height && y >= 0
	validWidth := x < maze.width && x >= 0
	if validHeight && validWidth {
		return true
	}
	return false
}

func (maze *maze) isGoal(x, y int) bool {
	if x == maze.goalX && y == maze.goalY {
		return true
	}
	return false
}

func solve(maze *maze, x int, y int) bool {

	maze.setPath(x, y)
	maze.solveVisited[x][y] = true

	if maze.isGoal(x, y) {
		return true
	}

	directions := shuffleDirections()

	for _, direction := range directions {
		switch direction {
		case up:
			newX := x
			newY := y + 1
			if maze.validNode(newX, newY) && !maze.solveVisited[newX][newY] && !maze.horizontalWall(newX, newY) {
				if solve(maze, newX, newY) {
					return true
				}
			}

		case down:
			newX := x
			newY := y - 1
			if maze.validNode(newX, newY) && !maze.solveVisited[newX][newY] && !maze.horizontalWall(newX, newY+1) {
				if solve(maze, newX, newY) {
					return true
				}
			}

		case left:
			newX := x - 1
			newY := y
			if maze.validNode(newX, newY) && !maze.solveVisited[newX][newY] && !maze.virticalWall(newX+1, newY) {
				if solve(maze, newX, newY) {
					return true
				}
			}

		case right:
			newX := x + 1
			newY := y
			if maze.validNode(newX, newY) && !maze.solveVisited[newX][newY] && !maze.virticalWall(newX, newY) {
				if solve(maze, newX, newY) {
					return true
				}
			}

		default:
			return false
		}
	}

	maze.unsetPath(x, y)
	return false
}

func (maze *maze) printSolution() {
	const hWallLineEnd = "+\n"
	const hWall = "+---"
	const vWallLineEnd = "|\n"

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

	if maze.isGoal(x, y) {

	}

	if maze.vWalls[x][y] {

		if maze.goalX == x && maze.goalY == y {
			return GoalVWall
		}

		if maze.path[x][y] {
			return vWallPath
		}

		return vWall
	}

	if maze.goalX == x && maze.goalY == y {
		return emptyGoal
	}

	if maze.path[x][y] {
		return emptyVWallPath
	}

	return emptyVWall
}

func (maze *maze) 


func (maze *maze) printWallSol(direction wallDirection, x int, y int) string {
	if direction == virtical {
		return maze.printViriticalWall(x, y)
	}
	return maze.printHorizontalWall(x, y)
}
