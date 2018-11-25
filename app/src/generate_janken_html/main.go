package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if isIllegalArgs(args) {
		os.Exit(1)
	}

	userHand := args[1]
	enemyHand := args[2]
	fmt.Println(userHand, enemyHand)

	// judge := judgeHand(userHand, enemyHand)
	// html := generateHTML(userHand, enemyHand, judge)
	// craeteTmpFile
	// mv tmp to location
}

func isHand(hand string) bool {
	switch hand {
	case "rock":
		return true
	case "paper":
		return true
	case "scissors":
		return true
	}
	return false
}

func isIllegalArgs(args []string) bool {
	if len(args) < 3 {
		return true
	}
	if !isHand(args[1]) {
		return true
	}
	if !isHand(args[2]) {
		return true
	}
	return false
}
