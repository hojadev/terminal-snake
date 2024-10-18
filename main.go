package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"math/rand"
	"time"
)

type GameBoard struct {
	Height int
	Width  int
	Cells  [][]string
	isOn   bool
}

type Position struct {
	X int
	Y int
}

type Food struct {
	Spawn    Position
	IsActive bool
}

type Snake struct {
	Body             []Position
	currentDirection Position
	Points           int
}

var snakeLetters = []Position{
	// S
	{2, 5}, {3, 5}, {4, 5}, {5, 5}, {6, 5}, {7, 5}, {8, 5}, {9, 5}, {10, 5}, {11, 5}, {12, 5}, // Parte superior
	{2, 6}, {3, 6}, {4, 6}, {5, 6}, {6, 6}, {7, 6}, {8, 6}, {9, 6}, {10, 6}, {11, 6}, {12, 6}, // Parte superior
	{2, 7}, {3, 7},
	{2, 8}, {3, 8},
	{2, 9}, {3, 9}, {4, 9}, {5, 9}, {6, 9}, {7, 9}, {8, 9}, {9, 9}, {10, 9}, {11, 9}, {12, 9}, // Parte superior
	{2, 10}, {3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}, {8, 10}, {9, 10}, {10, 10}, {11, 10}, {12, 10}, // Parte superior
	{11, 11}, {12, 11},
	{11, 12}, {12, 12},
	{2, 13}, {3, 13}, {4, 13}, {5, 13}, {6, 13}, {7, 13}, {8, 13}, {9, 13}, {10, 13}, {11, 13}, {12, 13}, // Parte superior
	{2, 14}, {3, 14}, {4, 14}, {5, 14}, {6, 14}, {7, 14}, {8, 14}, {9, 14}, {10, 14}, {11, 14}, {12, 14}, // Parte superior
	// N
	{14, 5}, {15, 5},
	{14, 6}, {15, 6},
	{14, 7}, {15, 7},
	{14, 8}, {15, 8},
	{14, 9}, {15, 9},
	{14, 10}, {15, 10},
	{14, 11}, {15, 11},
	{14, 12}, {15, 12},
	{14, 13}, {15, 13},
	{14, 14}, {15, 14},
	{16, 6}, {16, 5},
	{17, 7}, {17, 6},
	{18, 8}, {18, 7},
	{19, 9}, {19, 8},
	{20, 10}, {20, 9},
	{21, 11}, {21, 10},
	{22, 12}, {22, 11},
	{23, 13}, {23, 12},
	{24, 14}, {24, 13},
	{24, 5}, {25, 5},
	{24, 6}, {25, 6},
	{24, 7}, {25, 7},
	{24, 8}, {25, 8},
	{24, 9}, {25, 9},
	{24, 10}, {25, 10},
	{24, 11}, {25, 11},
	{24, 12}, {25, 12},
	{24, 13}, {25, 13},
	{24, 14}, {25, 14},
	//A
	{27, 5}, {28, 5}, {29, 5}, {30, 5}, {31, 5}, {32, 5}, {33, 5}, {34, 5}, {35, 5},
	{27, 6}, {28, 6}, {29, 6}, {30, 6}, {31, 6}, {32, 6}, {33, 6}, {34, 6}, {35, 6},
	{27, 7}, {28, 7},
	{27, 8}, {28, 8},
	{27, 9}, {28, 9},
	{27, 10}, {28, 10},
	{27, 11}, {28, 11},
	{27, 12}, {28, 12},
	{27, 13}, {28, 13},
	{27, 14}, {28, 14},
	{34, 7}, {35, 7},
	{34, 8}, {35, 8},
	{34, 9}, {35, 9},
	{34, 10}, {35, 10},
	{34, 11}, {35, 11},
	{34, 12}, {35, 12},
	{34, 13}, {35, 13},
	{34, 14}, {35, 14},
	{29, 9}, {30, 9}, {31, 9}, {32, 9}, {33, 9},
	{29, 10}, {30, 10}, {31, 10}, {32, 10}, {33, 10},
	//K

	{37, 5}, {38, 5},
	{37, 6}, {38, 6},
	{37, 7}, {38, 7},
	{37, 8}, {38, 8},
	{37, 9}, {38, 9},
	{37, 10}, {38, 10},
	{37, 11}, {38, 11},
	{37, 12}, {38, 12},
	{37, 13}, {38, 13},
	{37, 14}, {38, 14},
	{43, 5}, {44, 5},
	{43, 6}, {44, 6},
	{43, 7}, {44, 7},
	{42, 8}, {41, 8},
	{39, 9}, {40, 9},
	{40, 10}, {41, 10},
	{42, 11}, {43, 11},
	{43, 12}, {44, 12},
	{43, 13}, {44, 13},
	{43, 14}, {44, 14},
	//E
	{46, 5}, {47, 5}, {48, 5}, {49, 5}, {50, 5}, {51, 5}, {52, 5}, {53, 5}, {54, 5}, {55, 5},
	{46, 6}, {47, 6}, {48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{46, 7}, {47, 7},
	{46, 8}, {47, 8},
	{46, 9}, {47, 9}, {48, 9}, {49, 9}, {50, 9}, {51, 9}, {52, 9}, {53, 9}, {54, 9}, {55, 9}, // Parte superior
	{46, 10}, {46, 10}, {47, 10}, {48, 10}, {49, 10}, {50, 10}, {51, 10}, {52, 10}, {53, 10}, {54, 10}, {55, 10}, // Parte superior
	{46, 11}, {47, 11},
	{46, 12}, {47, 12},
	{46, 13}, {47, 13}, {48, 13}, {49, 13}, {50, 13}, {51, 13}, {52, 13}, {53, 13}, {54, 13}, {55, 13},
	{46, 14}, {47, 14}, {48, 14}, {49, 14}, {50, 14}, {51, 14}, {52, 14}, {53, 14}, {54, 14}, {55, 14},
}

func getValidNumber(w, h int) (xValid, yValid int) {
	xValid = rand.Intn(w - 1)
	yValid = rand.Intn(h - 1)

	if xValid == 0 {
		xValid = xValid + 1
	}

	if xValid == w {
		xValid = xValid - 1
	}

	if yValid == 0 {
		yValid = yValid + 1
	}

	if yValid == h {
		yValid = yValid - 1
	}
	return xValid, yValid
}
func (gb *GameBoard) grabInputFromMenu() {

	char, _, err := keyboard.GetKey()
	for {

		if err != nil {
			fmt.Println("Error al leer la tecla: ", err)
			return
		}

		if char != 'q' {
			gb.isOn = true
			break
		} else {
			gb.isOn = false
			break
		}

	}
}

func grabInputFromUser(s *Snake) {

	for {
		char, _, err := keyboard.GetKey()

		if err != nil {
			fmt.Println("Error al leer la tecla: ", err)
			return
		}

		if char != 0 {
			switch char {
			case 'w':
				if s.currentDirection.Y != 1 {
					s.currentDirection = Position{X: 0, Y: -1}
				}
			case 's':
				if s.currentDirection.Y != -1 {
					s.currentDirection = Position{X: 0, Y: 1}
				}
			case 'd':
				if s.currentDirection.X != -1 {
					s.currentDirection = Position{X: 1, Y: 0}
				}
			case 'a':
				if s.currentDirection.X != 1 {
					s.currentDirection = Position{X: -1, Y: 0}
				}
			}
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (gb GameBoard) Print() {
	for y := range gb.Cells {
		for x := range gb.Cells[y] {
			fmt.Print(gb.Cells[y][x])
		}
		fmt.Print("\n")
	}
}

func (gb *GameBoard) printScene(letters []Position) {

	for _, l := range letters {
		gb.Cells[l.Y][l.X] = "\033[32mS\033[0m"
	}
}

func (gb *GameBoard) updateState(s *Snake, f *Food) {

	for y := range gb.Cells {
		for x := range gb.Cells[y] {
			if gb.Cells[y][x] != "\033[31m#\033[0m" {
				gb.Cells[y][x] = " "
			}
		}
	}

	if f.IsActive == false {
		spawnX, spawnY := getValidNumber(gb.Width, gb.Height)
		f.Spawn.X = spawnX
		f.Spawn.Y = spawnY
		f.IsActive = true
	}

	gb.Cells[f.Spawn.Y][f.Spawn.X] = "\033[33m@\033[0m"

	for _, seg := range s.Body {
		if gb.Cells[seg.Y][seg.X] == "\033[31m#\033[0m" || gb.Cells[seg.Y][seg.X] == "\033[32m$\033[0m" {
			gb.isOn = false
			break
		}

		if gb.Cells[seg.Y][seg.X] == "\033[33m@\033[0m" {
			gb.Cells[seg.Y][seg.X] = "\033[32m$\033[0m"
			f.IsActive = false
			s.Points = s.Points + 1
			s.Body = append(s.Body, Position{X: s.Body[len(s.Body)-1].X, Y: s.Body[len(s.Body)-1].Y})
		}

		if gb.Cells[seg.Y][seg.X] == " " {
			gb.Cells[seg.Y][seg.X] = "\033[32m$\033[0m"
		}
	}
}

func gameLoop(gb *GameBoard, s *Snake, t *time.Ticker, f *Food) {

	for range t.C {

		clearScreen()

		gb.updateState(s, f)
		s.move()

		gb.Print()

		if s.Body[0].X >= gb.Width-1 {
			s.Body[0].X = 3
		}
		if !gb.isOn {
			fmt.Printf("Haz conseguido %d Puntos\n", s.Points)
			break
		}
	}
}

func (s *Snake) move() {

	newHead := Position{
		X: s.Body[0].X + s.currentDirection.X,
		Y: s.Body[0].Y + s.currentDirection.Y,
	}
	s.Body = append([]Position{newHead}, s.Body[:len(s.Body)-1]...)
}

func createSnake() Snake {
	s := Snake{
		Body: []Position{
			{X: 3, Y: 10},
			{X: 2, Y: 11},
			{X: 1, Y: 12},
		},
		currentDirection: Position{X: 1, Y: 0},
	}

	return s
}

func createBoard(h, w int) GameBoard {
	gb := GameBoard{
		Height: h,
		Width:  w,
		Cells:  make([][]string, h),
	}

	for y := range gb.Cells {
		gb.Cells[y] = make([]string, w)
		for x := range gb.Cells[y] {
			if (y == 0 || y == h-1) || (x == 0 || x == w-1) {
				gb.Cells[y][x] = "\033[31m#\033[0m"
			} else {
				gb.Cells[y][x] = " "
			}
		}
	}
	return gb
}

func main() {
	clearScreen()
	ticker := time.NewTicker(time.Duration(1000/15) * time.Millisecond)
	defer ticker.Stop()
	gb := createBoard(20, 60)

	if err := keyboard.Open(); err != nil {
		fmt.Println("Error al abrir el teclado:", err)
		return
	}
	defer keyboard.Close()

	for {
		gb.printScene(snakeLetters)
		gb.Print()
		fmt.Println("\nPresiona cualquier tecla para jugar o 'q' para salir")

		gb.grabInputFromMenu()
		if gb.isOn == true {
			gb = createBoard(20, 60)
			gb.isOn = true
			s := createSnake()
			f := Food{Spawn: Position{X: 5, Y: 5}, IsActive: true}

			go grabInputFromUser(&s)
			gameLoop(&gb, &s, ticker, &f)
			break

		} else {
			fmt.Println("Saliendo del juego...")
			break
		}
	}

}
