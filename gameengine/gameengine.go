package gameengine

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type LetterValidity string

const (
	PresentAndCorrectSpot LetterValidity = "correct"
	PresentButWrongSpot                  = "wrongspot"
	NotInAnySpot                         = "notpresent"
)

type Feedback [5]LetterValidity

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

	pretty := PrettyPrintFeedback(feedback)
	log.Printf("EvaluateSolution(%s vs %s) = %s\n", guess_word, reference_word, pretty)

	return feedback
}

const (
	GreenColor  = "\033[1;42m \033[0m "
	YellowColor = "\033[1;103m \033[0m "
	GreyColor   = "\033[1;47m \033[0m "
)

func PrettyPrintFeedback(feedback [5]LetterValidity) string {
	s := ""
	for _, f := range feedback {
		switch f {
		case PresentAndCorrectSpot:
			s += GreenColor
		case PresentButWrongSpot:
			s += YellowColor
		case NotInAnySpot:
			s += GreyColor
		}
	}
	return s
}

func ChooseRandomWord(dictionary_path string) string {
	file, err := os.Open(dictionary_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(lines))
	return lines[randomIndex]
}
