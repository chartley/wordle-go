// dev notes: https://docs.google.com/document/d/15DENBraewa74YN98-Mz700oZrgU-W1-WZnGUraNq_uU/edit

package main

import (
	"log"
	"wordle-go/gameengine"
	"wordle-go/playerenginelist"
	"wordle-go/playerenginetree"
)

const DictionaryPath = "./dictionary_len5.txt"

func solve_random_word_list() {
	// constant for "solved" feedback
	solvedFeedback := [5]gameengine.LetterValidity{
		gameengine.PresentAndCorrectSpot,
		gameengine.PresentAndCorrectSpot,
		gameengine.PresentAndCorrectSpot,
		gameengine.PresentAndCorrectSpot,
		gameengine.PresentAndCorrectSpot,
	}

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

func solve_random_word_tree() {
	// create the player
	player := playerenginetree.InitPlayer(DictionaryPath)

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
		guess_word := playerenginetree.Guess(player)
		feedback := gameengine.EvaluateSolution(target_word, guess_word)
		if feedback == solvedFeedback {
			log.Println("Solved in ", (i + 1))
			solved = true
		} else {
			playerenginetree.ProcessFeedback(player, guess_word, feedback)
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
