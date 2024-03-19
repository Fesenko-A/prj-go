package main

import (
	"fmt"
	"math/rand"
	"time"
)

var score int

func countdown(seconds int) {
	for seconds != 0 {
		fmt.Print(seconds, " ")
		time.Sleep(time.Second)
		seconds--
	}
	fmt.Println()
}

func generateNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func gameAdd(a, b int) {
	res := 0
	fmt.Print(a, " + ", b, " = ")
	fmt.Scanln(&res)
	if a+b == res {
		fmt.Println("True")
		score += 10
	} else {
		fmt.Println("False")
	}
}

func gameSub(a, b int) {
	res := 0
	fmt.Print(a, " - ", b, " = ")
	fmt.Scanln(&res)
	if a-b == res {
		fmt.Println("True")
		score += 10
	} else {
		fmt.Println("False")
	}
}

func gameMult(a, b int) {
	res := 0
	fmt.Print(a, " * ", b, " = ")
	fmt.Scanln(&res)
	if a*b == res {
		fmt.Println("True")
		score += 10
	} else {
		fmt.Println("False")
	}
}

func main() {
	fmt.Println("Welcome to my game!")
	countdown(5)

	start := time.Now()

	for score < 50 {
		a := generateNumber(10, 50)
		b := generateNumber(1, 10)

		gameType := generateNumber(1, 3)
		switch gameType {
		case 1:
			gameAdd(a, b)
		case 2:
			gameSub(a, b)
		case 3:
			gameMult(a, b)
		}
	}
	fmt.Println("You won! Congrats!")
	fmt.Println("Total time:", time.Since(start))

	fmt.Scanln()
}
