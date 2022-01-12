package gameengine

import "strings"

type LetterValidity string

const (
	PresentAndCorrectSpot LetterValidity = "correct"
	PresentButWrongSpot                  = "wrongspot"
	NotInAnySpot                         = "notpresent"
)

func EvaluateSolution(reference_word string, guess_word string) [5]LetterValidity {
	feedback := [5]LetterValidity{}
	for i := 0; i < len(guess_word); i += 1 {
		if guess_word[i] == reference_word[i] {
			feedback[i] = PresentAndCorrectSpot
		} else if strings.Contains(reference_word, string(guess_word[i])) {
			feedback[i] = PresentButWrongSpot
		} else {
			feedback[i] = NotInAnySpot
		}
	}
	return feedback
}
