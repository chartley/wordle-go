package main

import (
	"testing"
)

func TestEvaluateSolutioncorrect(t *testing.T) {
	guess_word := "hello"
	reference_word := "hello"
	feedback := evaluate_solution(reference_word, guess_word)
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
	feedback := evaluate_solution(reference_word, guess_word)
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
	feedback := evaluate_solution(reference_word, guess_word)
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
