// dev notes: https://docs.google.com/document/d/15DENBraewa74YN98-Mz700oZrgU-W1-WZnGUraNq_uU/edit

package main

import (
	"log"
	"wordle-go/gameengine"
	"wordle-go/playerenginelist"
)

const DictionaryPath = "./dictionary_len5.txt"

func solve_random_word_list() {
	// create the player
	player := playerenginelist.InitPlayer(DictionaryPath)

	// choose a random target word
	target_word := gameengine.ChooseRandomWord(DictionaryPath)
	log.Println("Target word", target_word)

	// constant for "solved" feedback
	solvedFeedback := [5]gameengine.LetterValidity{
		gameengine.PresentAndCorrectSpot,
		gameengine.PresentAndCorrectSpot,
		gameengine.PresentAndCorrectSpot,
		gameengine.PresentAndCorrectSpot,
		gameengine.PresentAndCorrectSpot,
	}

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

func main() {
	solve_random_word_list()
	// solve_random_word_tree()
	// alternatively solve a live puzzle, without knowing the word // alternatively solve a live puzzle, without knowing the word
	log.Println("Goodbye World!")
}
