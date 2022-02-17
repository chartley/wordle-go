package playerenginetree

import (
	"log"
	"wordle-go/gameengine"
	"wordle-go/playerstatetree"
)

type Player struct {
	solnspace    playerstatetree.Node // active solution space
	knownletters []rune               // letters we know exist but not where
	// letters we know are not in specific spots are pruned from tree
}

func InitPlayer(dictionary_path string) Player {
	player := Player{}
	player.solnspace = playerstatetree.BuildInitialTree(dictionary_path)
	player.knownletters = make([]rune, 0)
	return player
}

func Guess(player Player) string {
	// get frequency of letters remaining
	// score each word for how much it "represents" frequency of remaining letters
	// pick a guess word with the highest score
	return "guess"
}

func ProcessFeedback(player Player, guess_word string, feedback gameengine.Feedback) {
	for i, v := range feedback {
		char := guess_word[i]
		log.Println("ProcessFeedback", string(char), i, v)

		switch v {
		case gameengine.PresentAndCorrectSpot:
			// prune the tree for all words that have any other letter in this spot
		case gameengine.PresentButWrongSpot:
			// prune the tree for all words which have the letter in this spot
			// prune the tree for all words which don't have this letter
			// add to knownletters to inform guesses
		case gameengine.NotInAnySpot:
			// prune the tree for all words which have the letter
		}
	}
}
