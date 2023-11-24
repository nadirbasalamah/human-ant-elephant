package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var choiceIdx int
	var choices []string = []string{"Human", "Elephant", "Ant"}

	fmt.Println("Please choose:")
	for idx, c := range choices {
		fmt.Printf("%d) %s\n", idx+1, c)
	}

	fmt.Print("Inser your choice here: ")
	fmt.Scan(&choiceIdx)

	var isValid bool = validateInput(choiceIdx, len(choices))

	if isValid {
		choiceIdx--
	} else {
		fmt.Println("invalid input, please enter input from 1-3")
		return
	}

	playerChoice, opponentChoice := getChoices(choiceIdx, choices)

	var winner string = checkWinner(playerChoice, opponentChoice)

	fmt.Println(winner)
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
