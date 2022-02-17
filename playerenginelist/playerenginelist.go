package playerenginelist

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"wordle-go/gameengine"
)

type ScoredWord struct {
	word  string // the word
	score int    // scores, which we'll keep updating
}

type Player struct {
	possibleWords []*ScoredWord // words still eligible based on feedback
	// letters in correct positions (greens)
	// letters we know are NOT in specific spots (yellows)
	// letters we know DONT exist (greys)
}

func InitPlayer(dictionary_path string) Player {
	player := Player{}

	// open dictionary file
	words_file, err := os.Open(dictionary_path)
	if err != nil {
		log.Fatal(err)
	}
	defer words_file.Close()

	// init player state as full dictionary
	scanner := bufio.NewScanner(words_file)
	for scanner.Scan() {
		word := &ScoredWord{
			scanner.Text(),
			0,
		}
		player.possibleWords = append(player.possibleWords, word)
	}

	// init other state
	// player.possiblewords = playerstatetree.BuildInitialTree(dictionary_path)
	// player.knownletters = make([]rune, 0)

	return player
}

func Guess(player Player) string {
	log.Println("playerenginelist.Guess()")

	// count total character probabilities
	countPerLetter := map[rune]int{} // {char : count}
	for _, scoredWord := range player.possibleWords {
		for _, char := range scoredWord.word {
			countPerLetter[char] = countPerLetter[char] + 1
		}
	}
	x := 0
	for char, count := range countPerLetter {
		if x < 2 {
			log.Println("Count[", string(char), "] = ", count)
			x++
		} else {
			break
		}
	}

	// score each word
	for i, scoredWord := range player.possibleWords {
		score := 0
		seenChars := map[rune]bool{} // add score only once per char
		for _, char := range scoredWord.word {
			if _, charSeen := seenChars[char]; !charSeen {
				score = score + countPerLetter[char]
				seenChars[char] = true
			}
		}
		scoredWord.score = score
		if i < 3 {
			log.Println("Write Score[", scoredWord.word, "] = ", score)
		}
	}

	for i, scoredWord := range player.possibleWords {
		log.Println("Read Score[", scoredWord.word, "] = ", scoredWord.score)
		if i >= 2 {
			break
		}
	}

	// sort possibleWords by descending score to get ranked guesses
	sort.Slice(player.possibleWords, func(i, j int) bool {
		return player.possibleWords[i].score > player.possibleWords[j].score
	})
	for i, pw := range player.possibleWords {
		fmt.Printf("%s, %d\n", pw.word, pw.score)
		if i > 2 {
			break
		}
	}
	return player.possibleWords[0].word
}

func ProcessFeedback(player *Player, guess_word string, feedback gameengine.Feedback) {
	currWords := player.possibleWords
	newWords := []*ScoredWord{}
	for p, charFeedback := range feedback {
		char := guess_word[p]
		// log.Printf("ProcessFeedback(%s, p=%d, %s) on %d words\n", string(char), p, charFeedback, len(currWords))

		switch charFeedback {
		case gameengine.PresentAndCorrectSpot:
			// prune the tree for all words that have any other letter in this spot
			for _, scoredWord := range currWords {
				if scoredWord.word[p] == char {
					newWords = append(newWords, scoredWord)
				}
			}
		case gameengine.PresentButWrongSpot:
			// prune the tree of all words which have the letter in this spot
			for _, scoredWord := range currWords {
				if scoredWord.word[p] == char {
					// in the wrong spot, no longer eligible
				} else if strings.Index(scoredWord.word, string(char)) >= 0 {
					// has the char elsewhere though, so is eligible
					newWords = append(newWords, scoredWord)
				}
			}
			// TODO: add to knownletters to inform guesses?
		case gameengine.NotInAnySpot:
			// prune the tree for all words which have the letter at all
			for _, scoredWord := range currWords {
				if strings.Index(scoredWord.word, string(char)) < 0 {
					newWords = append(newWords, scoredWord)
				}
			}
		}

		// move that set to
		currWords = newWords
		newWords = []*ScoredWord{}
	}

	// log.Printf("ProcessFeedback() DONE on %d words\n", len(currWords))

	// write state back to player
	player.possibleWords = currWords
}
