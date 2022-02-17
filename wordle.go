// dev notes: https://docs.google.com/document/d/15DENBraewa74YN98-Mz700oZrgU-W1-WZnGUraNq_uU/edit

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"wordle-go/gameengine"
	"wordle-go/playerenginelist"
)

const DictionaryPath = "./dictionary_len5.txt"

// constant for "solved" feedback
var solvedFeedback = [...]gameengine.LetterValidity{
	gameengine.PresentAndCorrectSpot,
	gameengine.PresentAndCorrectSpot,
	gameengine.PresentAndCorrectSpot,
	gameengine.PresentAndCorrectSpot,
	gameengine.PresentAndCorrectSpot,
}

func solve_random_word_list() {
	// create the player
	player := playerenginelist.InitPlayer(DictionaryPath)

	// choose a random target word
	target_word := gameengine.ChooseRandomWord(DictionaryPath)
	log.Println("Target word", target_word)

	// take guesses until solved for out of tries
	solved := false
	for i := 0; i < 6; i += 1 {
		guess_word := playerenginelist.Guess(player)
		log.Printf("solve_random_word_list() i=%d, guess=%s", i, guess_word)
		feedback := gameengine.EvaluateSolution(target_word, guess_word)
		if feedback == solvedFeedback {
			log.Println("Solved in ", (i + 1))
			solved = true
			break
		} else {
			playerenginelist.ProcessFeedback(&player, guess_word, feedback)
		}
	}

	if !solved {
		log.Println("Failed to solve")
	}
}

func benchmark_random_word_list() {
	words := gameengine.ReadAllWords(DictionaryPath)
	attemps := make([]int, len(words))
	for j, target_word := range words {
		// create the player
		player := playerenginelist.InitPlayer(DictionaryPath)

		log.Println("Guessing", j, target_word)
		// take guesses until solved for out of tries
		solved := false
		for i := 0; i < 20; i += 1 {
			guess_word := playerenginelist.Guess(player)
			// log.Printf("solve_random_word_list() i=%d, guess=%s", i, guess_word)
			feedback := gameengine.EvaluateSolution(target_word, guess_word)
			if feedback == solvedFeedback {
				log.Println("Solved in ", (i + 1))
				attemps[j] = i + 1
				solved = true
				break
			} else {
				playerenginelist.ProcessFeedback(&player, guess_word, feedback)
			}
		}
		if !solved {
			log.Println("Failed to solve")
		}
	}

	min := 99999999
	max := -99999999
	sum := 0
	for _, a := range attemps {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
		sum += a
	}
	log.Println("Min", min, "Max", max, "Average: ", float64(sum)/float64(len(attemps)))
	log.Println()
}

func solve_interactive() {
	// create the player
	player := playerenginelist.InitPlayer(DictionaryPath)

	for {
		// tell user the guess, have them enter it into Wordle
		guess_word := playerenginelist.Guess(player)
		fmt.Printf("guess = %s input as [y=yes, s=somewhere, n=no] eg nnsns ; or q to quit\n", guess_word)
		fmt.Print("Enter feedback: ")

		// get the response back
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if len(text) < 5 {
			fmt.Println("Quitting...")
			break
		}

		// convert to gameengine.LetterValidity array
		feedback := [5]gameengine.LetterValidity{}
		for i := 0; i < 5; i += 1 {
			if text[i] == 'y' {
				feedback[i] = gameengine.PresentAndCorrectSpot
			} else if text[i] == 's' {
				feedback[i] = gameengine.PresentButWrongSpot
			} else if text[i] == 'n' {
				feedback[i] = gameengine.NotInAnySpot
			} else {
				log.Panicf("Invalid feedback char '%c'\n", text[i])
			}
		}

		// update player state
		playerenginelist.ProcessFeedback(&player, guess_word, feedback)
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "benchmark" {
		benchmark_random_word_list()
	} else if len(os.Args) > 1 && os.Args[1] == "interactive" {
		solve_interactive()
	} else {
		solve_random_word_list()
	}
}
