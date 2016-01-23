package quiz

import (
	"os"
	"testing"
)

//
// test data
//

var decomposeTestData = []struct {
	words    []string
	word     string
	expected []string
}{
	{
		[]string{"the", "longest", "day", "longestday"},
		"",
		[]string{},
	},
	{
		[]string{"the", "longest", "day", "longestday"},
		"longesttheday",
		[]string{"longest", "the", "day"},
	},
	{
		[]string{"de", "in", "st", "it", "ut", "io", "na", "li", "za", "ti", "on", "dei"},
		"deinstitutionalization",
		[]string{"de", "in", "st", "it", "ut", "io", "na", "li", "za", "ti", "on"},
	},
}

var testData = []struct {
	words           []string
	longestCompound string
	subwords        []string
}{
	{
		[]string{"the", "longest", "day", "longestday"},
		"longestday",
		[]string{"longest", "day"},
	},
	{
		[]string{"this", "is", "awesome"},
		"",
		nil,
	},
}

//
// Tests
//

// TestDecompose tests the Decompose call with a predefined set of inputs
func TestDecompose(t *testing.T) {
	for _, testCase := range decomposeTestData {
		subwords, err := Decompose(testCase.word, testCase.words)
		dieOnError(t, err)
		if !sameList(testCase.expected, subwords) {
			t.Fatalf("Expected subwords %v but got %v for '%s'", testCase.expected, subwords, testCase.word)
		}
	}
}

// TestLongestCompoundWord tests the LongestCompoundWord call with a predefined set of inputs
func TestLongestCompoundWord(t *testing.T) {
	for _, testCase := range testData {
		compound, subwords, err := LongestCompoundWord(testCase.words)
		dieOnError(t, err)
		verifyCompound(t, testCase.longestCompound, compound, testCase.subwords, subwords)
	}
}

// TestLongestCompoundWordList tests the LongestCompoundWord call with a user input file
func TestLongestCompoundWordList(t *testing.T) {
	input, err := os.Open("word.list")
	dieOnError(t, err)
	defer input.Close()
	words, err := ToWords(input)
	dieOnError(t, err)
	compound, subwords, err := LongestCompoundWord(words)
	dieOnError(t, err)
	expectedSubwords := []string{"de", "in", "st", "it", "ut", "io", "na", "li", "za", "ti", "on"}
	verifyCompound(t, "deinstitutionalization", compound, expectedSubwords, subwords)
}

//
// helper functions
//

// verifyCompound asserts the compound found is as expected
func verifyCompound(t *testing.T, expected, actual string, expectedSubwords, actualSubwords []string) {
	if expected != actual {
		t.Fatalf("Expected compound '%v' but got '%v'", expected, actual)
	}
	if !sameList(expectedSubwords, actualSubwords) {
		t.Fatalf("Expected subwords %v but got %v for compound '%v'", expectedSubwords, actualSubwords, expected)
	}
}

// dieOnError finish with a fatal on any non nil error
func dieOnError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// sameList returns true if both string lists are the same
func sameList(a, b []string) bool {
	if (a == nil && b != nil) || (b == nil && a != nil) {
		return false
	}
	if a == nil && b == nil {
		return true
	}
	// here both a and b are not nil...
	if len(a) != len(b) {
		return false
	}
	// here both a and b are the same size
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}
