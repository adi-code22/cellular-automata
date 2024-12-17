package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 700
	height = 400
	w      = 3
)

var (
	columns = width / w
	rows    = height / w
	board   [][]int
	next    [][]int
)

func init() {
	// Initialize the board with zeros
	board = make([][]int, columns)
	for i := range board {
		board[i] = make([]int, rows)
	}

	next = make([][]int, columns)
	for i := range next {
		next[i] = make([]int, rows)
	}

	// Initialize the grid
	initGrid()
}

func initGrid() {
	board[columns/4][0] = 1
	board[columns/2][0] = 1
	board[(columns*6)/8][0] = 1

	for j := 0; j < rows; j++ {
		for i := 0; i < columns; i++ {
			if i != 0 && j != 0 && i != columns-1 {
				XORCompare(i, j)
			}
		}
		for i2 := columns - 1; i2 > 0; i2-- {
			if i2 != 0 && j != 0 && i2 != columns-1 {
				XORCompare(i2, j)
			}
		}
	}
}

func XORCompare(i, j int) {
	UtL := (board[i-1][j-1] == 1)
	UuU := (board[i][j-1] == 1)
	UtR := (board[i+1][j-1] == 1)

	if (UtL && !(UuU || UtR)) || (!UtL && (UuU || UtR)) {
		board[i][j] = 1
	} else {
		board[i][j] = 0
	}
}

func update(screen *ebiten.Image) error {
	// Fill background with white
	screen.Fill(color.White)

	// Draw the grid cells
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			var c color.Color
			if board[i][j] == 1 {
				c = color.Black
			} else {
				c = color.White
			}

			// Create a rectangle (w-1) by (w-1) for each cell
			rectImage := image.NewRGBA(image.Rect(0, 0, w-1, w-1))
			for yi := 0; yi < w-1; yi++ {
				for xi := 0; xi < w-1; xi++ {
					rectImage.Set(xi, yi, c)
				}
			}

			// Draw the rectangle on the screen at the correct position
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*w), float64(j*w)) // Translate to the correct position
			screen.DrawImage(ebiten.NewImageFromImage(rectImage), op)
		}
	}

	return nil
}

type Game struct{}

// Update is called every frame to update the game state.
func (g *Game) Update() error {
	// No updates to game state needed, this is where you would add logic.
	return nil
}

// Draw is called every frame to render the game state.
func (g *Game) Draw(screen *ebiten.Image) {
	update(screen)
}

// Layout is required by ebiten.Game and specifies the screen size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	// Create a new Ebiten game window
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Go Grid Simulation")

	// Start the game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
