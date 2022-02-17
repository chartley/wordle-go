package playerenginelist

import (
	"testing"
	"wordle-go/gameengine"
)

func TestInitPlayer(t *testing.T) {
	InitPlayer("../dictionary_len5.txt")
}

func TestInitPlayerGuess(t *testing.T) {
	player := InitPlayer("../dictionary_len5.txt")
	Guess(player)
}

// guess = weeee, actualWord = woooo, dictionary = [woooo, gaaaa]
func TestProcessFeedbackSinglePresentAndCorrectSpot(t *testing.T) {
	player := Player{}
	player.possibleWords = append(player.possibleWords, &ScoredWord{"woooo", 0})
	player.possibleWords = append(player.possibleWords, &ScoredWord{"gaaaa", 0})
	guess_word := "weeee"
	feedback := [5]gameengine.LetterValidity{
		gameengine.PresentAndCorrectSpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
	}
	ProcessFeedback(&player, guess_word, feedback)
	if len(player.possibleWords) != 1 {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
	if player.possibleWords[0].word != "woooo" {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
}

func TestProcessFeedbackSinglePresentButWrongSpot1(t *testing.T) {
	player := Player{}
	player.possibleWords = append(player.possibleWords, &ScoredWord{"woooo", 0})
	player.possibleWords = append(player.possibleWords, &ScoredWord{"gaaaa", 0})
	guess_word := "xxxxw"
	feedback := [5]gameengine.LetterValidity{
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.PresentButWrongSpot,
	}
	ProcessFeedback(&player, guess_word, feedback)
	if len(player.possibleWords) != 1 {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
	if player.possibleWords[0].word != "woooo" {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
}

func TestProcessFeedbackSinglePresentButWrongSpot2(t *testing.T) {
	player := Player{}
	player.possibleWords = append(player.possibleWords, &ScoredWord{"woooo", 0})
	player.possibleWords = append(player.possibleWords, &ScoredWord{"gaaaa", 0})
	guess_word := "xwxxx"
	feedback := [5]gameengine.LetterValidity{
		gameengine.NotInAnySpot,
		gameengine.PresentButWrongSpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
	}
	ProcessFeedback(&player, guess_word, feedback)
	if len(player.possibleWords) != 1 {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
	if player.possibleWords[0].word != "woooo" {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
}
func TestProcessFeedbackAllNotInAnySpotMiss(t *testing.T) {
	player := Player{}
	player.possibleWords = append(player.possibleWords, &ScoredWord{"aaaaa", 0})
	player.possibleWords = append(player.possibleWords, &ScoredWord{"bbbbb", 0})
	guess_word := "ccccc"
	feedback := [5]gameengine.LetterValidity{
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
	}
	ProcessFeedback(&player, guess_word, feedback)
	if len(player.possibleWords) != 2 {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
}

func TestProcessFeedbackAllNotInAnySpotHit(t *testing.T) {
	player := Player{}
	player.possibleWords = append(player.possibleWords, &ScoredWord{"aaaaa", 0})
	player.possibleWords = append(player.possibleWords, &ScoredWord{"bbbbb", 0})
	guess_word := "bbbbb"
	feedback := [5]gameengine.LetterValidity{
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
		gameengine.NotInAnySpot,
	}
	ProcessFeedback(&player, guess_word, feedback)
	if len(player.possibleWords) != 1 {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
	if player.possibleWords[0].word != "aaaaa" {
		t.Fatalf("len(player.possibleWords) = %d", len(player.possibleWords))
	}
}
