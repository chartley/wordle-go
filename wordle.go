// dev notes: https://docs.google.com/document/d/15DENBraewa74YN98-Mz700oZrgU-W1-WZnGUraNq_uU/edit

package main

import (
	"log"
	"wordle-go/gameengine"
	"wordle-go/playerengine"
)

func solve_random_word() {
	// create the player
	dictionary_path := "./dictionary_len5.txt"
	player := playerengine.InitPlayer(dictionary_path)

	// choose a random target word
	target_word := gameengine.ChooseRandomWord(dictionary_path)
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
		guess_word := playerengine.Guess(player)
		feedback := gameengine.EvaluateSolution(target_word, guess_word)
		if feedback == solvedFeedback {
			log.Println("Solved in ", (i + 1))
			solved = true
		} else {
			playerengine.ProcessFeedback(player, feedback)
		}
	}

	if !solved {
		log.Println("Failed to solve")
	}
}

func main() {
	solve_random_word() // alternatively solve a live puzzle, without knowing the word
	log.Println("Goodbye World!")
}
