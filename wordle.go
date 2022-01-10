// dev notes: https://docs.google.com/document/d/15DENBraewa74YN98-Mz700oZrgU-W1-WZnGUraNq_uU/edit

package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type LetterValidity string

const (
	PresentAndCorrectSpot LetterValidity = "correct"
	PresentButWrongSpot                  = "wrongspot"
	NotInAnySpot                         = "notpresent"
)

type node struct {
	children map[rune]node
}

func init_node() node {
	n := node{}
	n.children = map[rune]node{}
	return n
}

func build_initial_tree(dictionary_path string) node {
	words_file, err := os.Open(dictionary_path)
	if err != nil {
		log.Fatal(err)
	}
	defer words_file.Close()

	// add each word to the solution space
	scanner := bufio.NewScanner(words_file)
	root := init_node()
	for scanner.Scan() {
		add_word(root, scanner.Text())
	}

	dump_tree(root)

	return root
}

func add_word(root node, word string) {
	node_here := root
	for word_i := 0; word_i < len(word); word_i++ {
		letter := rune(word[word_i])
		child_node, found := node_here.children[letter]
		if found {
			// found, build down the tree
			node_here = child_node
		} else {
			// not found; add node then iterate
			node_here.children[letter] = init_node()
			node_here = node_here.children[letter]
		}
	}
}

func dump_tree(root node) {
	nodes_level := []node{root}
	n_nodes := 1
	n_levels := 1
	nodes_next := []node{}
	for len(nodes_level) > 0 {
		for i := 0; i < len(nodes_level); i += 1 {
			for _, child_node := range nodes_level[i].children {
				nodes_next = append(nodes_next, child_node)
				n_nodes += 1
				// log.Println("dump_tree() i", i, "nodes_level, ", len(nodes_level), "nodes_next", len(nodes_next), "char", string(char))
			}
		}

		// iterate through next level
		nodes_level = nodes_next
		nodes_next = []node{}
		n_levels += 1
	}

	log.Println("dump_tree()", n_nodes, "nodes, ", n_levels, "levels")
}

func evaluate_solution(reference_word string, guess_word string) [5]LetterValidity {
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

func main() {
	// sanity check - test word is in dictionary :) wordle is uk spelling

	// create problem space from full dictionary
	//space :=
	build_initial_tree("./dictionary_len5.txt")

	log.Println("Goodbye World!")
}
