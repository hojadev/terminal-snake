package main

import (
	"fmt"
	"testing"
)

func TestRandNumbersXY(t *testing.T) {
	for range 100 {
		x, y := getValidNumber(20, 30)
		if x == 0 || y == 0 {
			fmt.Println("FAIL TEST: ", x, y)
		}
		//		fmt.Printf("(x: %d,y: %d)\n", x, y)
	}

}

func testSnakeLetters(t *testing.T) {
	for xy := range snakeLetters {
		fmt.Println(xy)
	}
}
