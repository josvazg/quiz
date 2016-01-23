package quiz

import (
	"bytes"
	"fmt"
	"unicode"
)

const (
	NO_LETTER       = 0
	MaxAsciiLetters = 27
)

// Trie is the root or a node in trie structure
type Trie struct {
	letter   rune
	terminal bool
	childs   [MaxAsciiLetters]*Trie
}

// Add a word to the Trie
func (root *Trie) Add(word string) error {
	node := root
	for _, letter := range word {
		idx := asciiDict(letter)
		if idx == NO_LETTER || idx < 0 || idx >= MaxAsciiLetters {
			return fmt.Errorf("Can't add word '%s' with invalid letter %v", word, letter)
		}
		if node.childs[idx] == nil {
			node.childs[idx] = &Trie{letter: unicode.ToLower(letter)}
		}
		node = node.childs[idx]
	}
	node.terminal = true
	return nil
}

// Prefixes returns all valid prefixes found for the given text,
// including text itself, if it happens to be a found within the trie
func (root *Trie) Prefixes(text string) ([]string, error) {
	prefix := bytes.NewBufferString("")
	prefixes := []string{}
	node := root
	for _, letter := range text {
		idx := asciiDict(letter)
		if idx == NO_LETTER || idx < 0 || idx >= MaxAsciiLetters {
			return nil, fmt.Errorf("text '%s' has invalid letter %v", text, letter)
		}
		if node.childs[idx] == nil {
			return prefixes, nil
		}
		node = node.childs[idx]
		prefix.WriteRune(node.letter)
		if node.terminal {
			prefixes = append(prefixes, prefix.String())
		}
	}
	return prefixes, nil
}

// dictAscii maps Ascii only runes to a lower case index 1-26, or returns 0 on NO_LETTER
func asciiDict(letter rune) int {
	if letter >= 'a' && letter <= 'z' {
		return (int)(letter - 'a' + 1)
	}
	if letter >= 'A' && letter <= 'Z' {
		return (int)(letter - 'A' + 1)
	}
	return NO_LETTER
}
