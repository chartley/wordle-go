package playerengine

import (
	"wordle-go/gameengine"
	"wordle-go/playerstate"
)

type Player struct {
	solnspace playerstate.Node // active solution space
	// letters we know exist but not where
	// letters we know are not in specific spots <- pruned from tree
}

func InitPlayer(dictionary_path string) Player {
	player := Player{}
	player.solnspace = playerstate.BuildInitialTree(dictionary_path)
	// player.
	return player
}

func Guess(player Player) string {
	return "guess"
}

func ProcessFeedback(player Player, feedback gameengine.Feedback) {

}
