package playerstate

import (
	"bufio"
	"log"
	"os"
	"reflect"
)

type Node struct {
	children map[rune]Node
}

func init_node() Node {
	n := Node{}
	n.children = map[rune]Node{}
	return n
}

func BuildInitialTree(dictionary_path string) Node {
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

func BuildTreeFromWords(words []string) Node {
	root := init_node()
	for _, word := range words {
		add_word(root, word)
	}

	dump_tree(root)
	return root
}

func add_word(root Node, word string) {
	node_here := root
	for word_i := 0; word_i < len(word); word_i++ {
		letter := rune(word[word_i])
		child_node, found := node_here.children[letter]
		if found {
			// found, build down the tree
			node_here = child_node
		} else {
			// not found; add Node then iterate
			node_here.children[letter] = init_node()
			node_here = node_here.children[letter]
		}
	}
}

func dump_tree(root Node) {
	nodes_level := []Node{root}
	n_nodes := 1
	n_levels := 1
	nodes_next := []Node{}
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
		nodes_next = []Node{}
		n_levels += 1
	}

	// log.Println("dump_tree()", n_nodes, "nodes, ", n_levels, "levels")
}

func get_all_words(root Node) []string {
	all_words := []string{}
	for char, child_node := range root.children {
		words := get_words_recurse(child_node, string(char))
		all_words = append(all_words, words...)
	}
	return all_words
}

func get_words_recurse(node Node, prefix string) []string {
	// log.Println("get_words_recurse() prefix", prefix)
	all_words := []string{}
	for char, child_node := range node.children {
		keys := reflect.ValueOf(child_node.children).MapKeys()
		if len(keys) > 0 {
			words := get_words_recurse(child_node, prefix+string(char))
			all_words = append(all_words, words...)
		} else {
			all_words = append(all_words, prefix+string(char))
		}
	}
	return all_words
}
