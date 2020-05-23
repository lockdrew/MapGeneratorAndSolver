package main

import (
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	visted bool
	wall bool
}

type maze struct {
	mazeMap [][]node
	solved  bool
	height int 
	width int 
}

type Direction int;

const (
	up Direction = iota
	down
	left
	right  
) 

var Directions = []Direction { 
	up, 
	down, 
	left,
	right,
}


func main() {
	rand.Seed(time.Now().UnixNano())

	var maze = initMaze(3, 3)

	maze.printMaze()

	//Randomly Selected Start
	recursiveBackTrack(maze, 1, 1)

	
}

//Create the maze map and set initial values 
func initMaze(height, width int) *maze {
	var mazeMap = make([][]node, height)
	var maze = &maze{mazeMap, false, height, width}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			var newNode = &node{false, true}
			maze.mazeMap[h] = append(maze.mazeMap[h], *newNode)
		}
	}

	return maze
}

//Print Maze 
func (mazeToPrint *maze) printMaze() { 
	for h := 0; h < mazeToPrint.height; h++ {
		fmt.Println("")
		for w := 0; w < mazeToPrint.width; w++ { 
			var nodeToPrint = mazeToPrint.mazeMap[w][h]
			nodeToPrint.printNode()
		}
	}
	fmt.Println("")
}

//Prints either "#" for wall or " " if not wall
func (nodeToPrint *node) printNode() {
	if nodeToPrint.wall { 
		print("#")
	}
	print(" ")
}


//Check if node is Inside the Map
func (mazeToValidate *maze) validNode(x int, y int) bool { 
	var validHeight = y < mazeToValidate.height && y >= 0 
	var validWidth = x < mazeToValidate.width && x >= 0
	if ( validHeight && validWidth) { 
			return true 
	} 
	return false
}

func (maze *maze) visted(x int, y int) bool { 
	return maze.mazeMap[x][y].visted
}

func (maze *maze) setVisted(x int, y int) {
	maze.mazeMap[x][y].visted = true
}

func (maze *maze) removeWall(x int, y int) { 
	maze.mazeMap[x][y].wall = false
}

func recursiveBackTrack(maze *maze, x int, y int) { 

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Directions), func(i, j int) { 
		Directions[i] , Directions[j] = Directions[j] , Directions[i]
	})

	//Go through all child nodes
	for _ , direction := range Directions { 
		maze.printMaze()
		x, y := navigate(direction, x, y)
		fmt.Printf("looking at node %d, %d\n", x , y)
		// fmt.Printf("valid node %t\n", maze.validNode(x, y))
		// fmt.Printf("visted node %t\n", maze.visted(x, y))

		time.Sleep(2 * time.Second)

		if maze.validNode(x, y) && !maze.visted(x, y) { 
			fmt.Printf("removed %d, %d\n",x,y)
			maze.setVisted(x, y)
			maze.removeWall(x, y)
			recursiveBackTrack(maze, x, y) 
		}
	}

	//No Path Forward Found - BackTrack
	maze.printMaze()
	return 
}

func navigate(directionToMove Direction, x int, y int) (int, int) { 
	switch directionToMove { 
	case up: 
		return x, y+1
	case down: 
		return x, y-1
	case left: 
		return x-1, y
	case right: 
		return x+1, y
	default: 
	    return 0,0
	}
}