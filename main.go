package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CellContent int

const (
	SCREEN_WIDTH  int32 = 640
	SCREEN_HEIGHT int32 = 480
	CELL_SIZE     int32 = 20

	FIRE_SPREAD_CHANGE int = 30
	BURN_OUT_CHANGE    int = 70
	TREE_REGROW_CHANCE int = 10
)

const (
	Empty     CellContent = iota
	Tree      CellContent = iota
	Fire      CellContent = iota
	BurnedOut CellContent = iota
)

var (
	rows    int32    = SCREEN_HEIGHT / CELL_SIZE
	columns int32    = SCREEN_WIDTH / CELL_SIZE
	cells   [][]Cell = [][]Cell{}
)

type Cell struct {
	x       int32 // X Position on the grid, not on the screen pixel
	y       int32 // Y Positio on the grid, not on the screen pixeln
	content CellContent
	color   rl.Color
}

func NewCell(x int32, y int32, content CellContent) Cell {
	cell := Cell{x: x, y: y, content: content, color: rl.White}

	return cell
}

func (cell Cell) Draw() {
	xPos := cell.x * CELL_SIZE
	yPos := cell.y * CELL_SIZE

	color := rl.Black
	switch cell.content {
	case Tree:
		color = rl.Green
	case Fire:
		color = rl.Red
	case BurnedOut:
		color = rl.Black
	}

	rl.DrawRectangle(xPos, yPos, CELL_SIZE, CELL_SIZE, color)
}

func (cell Cell) GetNeighbours() []Cell {
	x := cell.x
	y := cell.y

	neighbours := []Cell{}
	for xIndex := -1; xIndex <= 1; xIndex++ {
		for yIndex := -1; yIndex <= 1; yIndex++ {
			if xIndex == 0 && yIndex == 0 { // Ignore the cell itself - it is not a neighbour of itself
				continue
			}

			neighbourX := x + int32(xIndex)
			neighbourY := y + int32(yIndex)

			if neighbourX >= 0 && neighbourX < int32(len(cells)) && neighbourY >= 0 && neighbourY < int32(len(cells[0])) {
				neighbours = append(neighbours, cells[neighbourX][neighbourY])
			}
		}
	}

	return neighbours
}

func FireShouldSpread() bool {
	randomNumber := rand.Intn(100) + 1
	if randomNumber <= FIRE_SPREAD_CHANGE {
		return true
	}

	return false
}

func FireShouldBurnOut() bool {
	randomNumber := rand.Intn(100) + 1
	if randomNumber <= BURN_OUT_CHANGE {
		return true
	}

	return false
}

func UpdateFire() {
	newCells := make([][]Cell, len(cells))
	for x := range cells {
		newCells[x] = make([]Cell, len(cells[x]))
		copy(newCells[x], cells[x])
	}

	for x := range columns {
		for y := range rows {
			for _, nCell := range cells[x][y].GetNeighbours() {
				if nCell.content == Fire {
					if FireShouldSpread() {
						newCells[x][y].content = Fire
					}

					if FireShouldBurnOut() {
						newCells[x][y].content = BurnedOut
					}
				}
			}
		}
	}

	cells = newCells
}

func ShouldTreeRegrow() bool {
	randomNumber := rand.Intn(100) + 1
	if randomNumber <= TREE_REGROW_CHANCE {
		return true
	}

	return false
}

func UpdateBurnedOut() {
	newCells := make([][]Cell, len(cells))
	for x := range cells {
		newCells[x] = make([]Cell, len(cells[x]))
		copy(newCells[x], cells[x])
	}

	for x := range columns {
		for y := range rows {
			if newCells[x][y].content == BurnedOut {
				if ShouldTreeRegrow() {
					newCells[x][y].content = Tree
				}
			}
		}
	}

	cells = newCells
}

func Init() {
	for col := range columns {
		colCells := []Cell{}
		for row := range rows {
			content := Empty
			randomNumber := rand.Intn(100) + 1
			if randomNumber >= 1 && randomNumber <= 99 {
				content = Tree
			} else if randomNumber > 99 {
				content = Fire
			}

			cell := NewCell(int32(col), int32(row), content)
			colCells = append(colCells, cell)
		}

		cells = append(cells, colCells)
	}
}

func Update() {
	UpdateFire()
	UpdateBurnedOut()
}

func Draw() {
	for _, cellSlice := range cells {
		for _, cell := range cellSlice {
			cell.Draw()
		}
	}
}

func main() {
	Init()

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Forest Fire")
	defer rl.CloseWindow()

	rl.SetTargetFPS(5)

	for !rl.WindowShouldClose() {
		Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		Draw()
		rl.EndDrawing()
	}
}
