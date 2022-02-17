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
	return feedback
}

func ChooseRandomWord(dictionary_path string) string {
	lines := ReadAllWords(dictionary_path)

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(lines))
	return lines[randomIndex]
}

func ReadAllWords(dictionary_path string) []string {
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

	return lines
}
