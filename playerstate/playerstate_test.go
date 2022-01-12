package playerstate

import (
	"reflect"
	"sort"
	"testing"
)

func TestBuildTreeGetAllWordsSeparate(t *testing.T) {
	input_words := []string{
		"hello",
		"there",
		"world",
	}
	root := BuildTreeFromWords(input_words)
	dumped_words := get_all_words(root)
	sort.Strings(dumped_words)
	if !reflect.DeepEqual(input_words, dumped_words) {
		t.Fatal("dumped_words", dumped_words)
	}
}

func TestBuildTreeGetAllWordsPrefixed(t *testing.T) {
	input_words := []string{
		"hello",
		"hells",
		"helps",
	}
	root := BuildTreeFromWords(input_words)
	dumped_words := get_all_words(root)
	sort.Strings(dumped_words)
	if !reflect.DeepEqual(input_words, dumped_words) {
		t.Fatal("dumped_words", dumped_words)
	}
}
