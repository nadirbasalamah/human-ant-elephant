package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	var playerChoiceIdx int
	var continuePlaying string
	var choices []string = []string{"Human", "Elephant", "Ant"}

	for {
		for {
			choiceIdx, isValid := mainMenu(choices)

			if isValid {
				playerChoiceIdx = choiceIdx
				break
			}
		}

		playerChoice, opponentChoice := getChoices(playerChoiceIdx, choices)

		var winner string = checkWinner(playerChoice, opponentChoice)

		fmt.Println(winner)

		fmt.Print("Continue playing? (y/n): ")
		fmt.Scan(&continuePlaying)

		if continuePlaying == "n" {
			fmt.Println("Bye...")
			break
		} else if continuePlaying != "y" && continuePlaying != "n" {
			fmt.Println("please enter n to quit the game")
		} else {
			clearTerminal()
		}
	}
}

func mainMenu(choices []string) (int, bool) {
	var choiceIdx int

	fmt.Println("Please choose:")
	for idx, c := range choices {
		fmt.Printf("%d) %s\n", idx+1, c)
	}

	fmt.Print("Insert your choice here: ")
	fmt.Scan(&choiceIdx)

	var isValid bool = validateInput(choiceIdx, len(choices))

	if isValid {
		choiceIdx--
		return choiceIdx, true
	} else {
		clearTerminal()
		fmt.Println("invalid input, please enter input from 1-3")
		return -1, false
	}
}

func validateInput(choiceIdx, numOfChoices int) bool {
	var isInvalid bool = choiceIdx <= 0 || choiceIdx >= numOfChoices+1

	if isInvalid {
		return false
	}

	return true
}

func getChoices(playerChoiceIdx int, choices []string) (string, string) {
	var opponentChoiceIdx int = rand.Intn(len(choices))

	var playerChoice string = choices[playerChoiceIdx]
	var opponentChoice string = choices[opponentChoiceIdx]

	fmt.Println("Player choose: ", playerChoice)
	fmt.Println("Opponent choose: ", opponentChoice)

	return playerChoice, opponentChoice
}

func checkWinner(player, opponent string) string {
	var isPlayerWin bool = (player == "Human" && opponent == "Ant") ||
		(player == "Elephant" && opponent == "Human") ||
		(player == "Ant" && opponent == "Elephant")

	var isOpponentWin bool = (player == "Ant" && opponent == "Human") ||
		(player == "Human" && opponent == "Elephant") ||
		(player == "Elephant" && opponent == "Ant")

	if player == opponent {
		return "tie!"
	} else if isPlayerWin {
		return "you win!"
	} else if isOpponentWin {
		return "you lose!"
	} else {
		return "tie!"
	}
}

func clearTerminal() {
	clearTerminalFunc, ok := clear[runtime.GOOS]
	if ok {
		clearTerminalFunc()
	} else {
		panic("unsupported platform. clear terminal failed")
	}
}
