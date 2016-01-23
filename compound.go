package quiz

import (
	"bufio"
	"io"
)

// ToWords extracts all words found from Reader r
func ToWords(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	words := []string{}
	// Set the split function for the scanning operation to words
	scanner.Split(bufio.ScanWords)
	// get the words, report any error or return them
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			words = append(words, s)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

// LongestCompoundWord return the longest compound word found in the words list
//
// Process:
// 1) Put all words on a Trie
// 2) Loop through all words decomposing them
// 3) Record and return the longest compound found amongst all decompositions
func LongestCompoundWord(words []string) (string, []string, error) {
	longestCompound := ""
	var longestSubwords []string = nil
	skipped := 0
	trie, err := toTrie(words)
	if err != nil {
		return "", nil, err
	}
	for _, word := range words {
		if len(word) >= len(longestSubwords) {
			subwords, err := decompose(word, trie)
			if err != nil {
				return "", nil, err
			}
			if len(subwords) > 1 && len(subwords) > len(longestSubwords) {
				longestCompound = word
				longestSubwords = subwords
			}
		} else {
			skipped++
		}
	}
	return longestCompound, longestSubwords, nil
}

// Decompose returns the longest possible decomposition of word into subwords found in words,
// that includes the whole word itself only if it is found in words
//
// (see decompose() comments for algorithm details)
func Decompose(word string, words []string) ([]string, error) {
	trie, err := toTrie(words)
	if err != nil {
		return nil, err
	}
	return decompose(word, trie)
}

//
// Private functions
//

// toTrie creates a trie from the given set of words
func toTrie(words []string) (*Trie, error) {
	trie := &Trie{}
	for _, word := range words {
		if err := trie.Add(word); err != nil {
			return nil, err
		}
	}
	return trie, nil
}

// decompose tries to decompose word into the longest sequence of words found in the given trie
//
// Steps:
// FOR each prefix of word in found in trie...
//   IF prefix is word AND we have no subwords recorded, record subwords = [prefix]
//   Compute remainder suffix
//   decompose suffix into suffixes
//   IF any suffixes are returned AND their are more or equal the recorded subwords length so far...
//     record subwords = [prefix, {suffixes}] as result
// return subwords
//
// Invariants and validations:
// 1) Empty word returns zero length array []
// 2) Words with no prefixes ARE not present in the trie AND also return []
// 3) If the only prefix of word in the trie happens to be word, [word] is returned
// 4) subwords is the longest decomposition found so far
func decompose(word string, trie *Trie) ([]string, error) {
	subwords := []string{}
	prefixes, err := trie.Prefixes(word)
	if err != nil {
		return nil, err
	}
	for _, prefix := range prefixes {
		if prefix == word && len(subwords) == 0 {
			subwords = []string{prefix}
			continue
		}
		suffix := word[len(prefix):]
		suffixes, err := decompose(suffix, trie)
		if err != nil {
			return nil, err
		}
		if len(suffixes) > 0 && len(suffixes) >= len(subwords) {
			subwords = append([]string{prefix}, suffixes...)
		}
	}
	return subwords, nil
}
