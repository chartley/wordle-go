package gameengine

import (
	"testing"
)

func TestEvaluateSolutioncorrect(t *testing.T) {
	guess_word := "hello"
	reference_word := "hello"
	feedback := EvaluateSolution(reference_word, guess_word)
	expected_feedback := [5]LetterValidity{
		PresentAndCorrectSpot,
		PresentAndCorrectSpot,
		PresentAndCorrectSpot,
		PresentAndCorrectSpot,
		PresentAndCorrectSpot,
	}
	if feedback != expected_feedback {
		t.Fatal(feedback)
	}
}

func TestEvaluateSolutionAllWrong(t *testing.T) {
	guess_word := "hello"
	reference_word := "grain"
	feedback := EvaluateSolution(reference_word, guess_word)
	expected_feedback := [5]LetterValidity{
		NotInAnySpot,
		NotInAnySpot,
		NotInAnySpot,
		NotInAnySpot,
		NotInAnySpot,
	}
	if feedback != expected_feedback {
		t.Fatal(feedback)
	}
}

func TestEvaluateSolutionSingleMisplacedLetter(t *testing.T) {
	guess_word := "hello"
	reference_word := "chain"
	feedback := EvaluateSolution(reference_word, guess_word)
	expected_feedback := [5]LetterValidity{
		PresentButWrongSpot,
		NotInAnySpot,
		NotInAnySpot,
		NotInAnySpot,
		NotInAnySpot,
	}
	if feedback != expected_feedback {
		t.Fatal(feedback)
	}
}
