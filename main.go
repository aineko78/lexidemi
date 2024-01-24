package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
)

// Standard helper for error checks
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func unfairGram(probabilities map[string]int, total int) string {
	//To extract n-grams with unequal probabilities we need a function to simulate an unfair dice with total sides
	//We use total because converting a random number in the unit interval to total is simpler than the reverse
	//and avoids floats that would not add up exactly to 1

	// Generate a random number between 0 and total. We are content with taking the integer part since rounding would change our pseudorandom number
	randomNumber := int(rand.Float64() * float64(total))

	// Determine the outcome based on probabilities
	cumulativeProb := 0
	for i, prob := range probabilities {
		cumulativeProb += prob
		if randomNumber < cumulativeProb {
			return i
		}
	}
	return ""
}

func generateWord(dict []map[string]map[string]int) string {

	sum := 0
	firstLetter := ""

	//As first letter we choose only among those that are initials, so we are a bit less non-stationary (to be fully
	//stationary we would need to consider the last letter of the previous word)
	for _, value := range dict[1][" "] {
		sum += value
	}

	firstLetter = unfairGram(dict[1][" "], sum)

	word := []rune(" " + firstLetter)
	b := firstLetter
	c := string(word)
	//Other letters after the second until space
	for ok := true; ok; ok = (b != " " && b != "" && len(word) < 15) { //Lentgh is necessary if there are no spaces in the parsed file.
		sum2 := 0
		for _, value := range dict[2][c] {
			sum2 += value
		}
		b = unfairGram(dict[2][c], sum2)
		word = append(word, []rune(b)...) //Even if only one rune is appended, we still need "..."
		c = string(word[len(word)-2:])
	}

	return string(word)
}

func main() {
	var numWords int
	var file string
	var help bool

	flag.IntVar(&numWords, "n", 3, "Number of words to generate")
	flag.StringVar(&file, "f", "", "File to parse")
	//	flag.StringVar(&save, "s", "", "File to save the frequencies")
	flag.BoolVar(&help, "help", false, "Help")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	//Even if the file were to contain the entire dictionary, no way would it exceed memory available in any modern computer, so we just load it whole
	dat, err := os.ReadFile(file)
	check(err)

	//Indexing a slice accesses individual bytes, not characters, hence we use slices of runes. This is essential because
	//a language may include any phonemes, e. i. long vowels.
	//Since a key will include 1,2 or 3 chars, slice of runes to compute frequencies but keys get converted back to strings for maps.

	arr := make([]rune, len(string(dat)))               //We know the length, so we preallocate for performance.
	re := regexp.MustCompile(`\r?\n`)                   //We have to use regex to deal with newlines in all formats
	arr = []rune(re.ReplaceAllString(string(dat), " ")) // Go is strongly typed, hence we have to convert from byte slice to string to rune
	s, a, b := []rune(""), []rune(""), []rune("")
	dict := make([]map[string]map[string]int, 3) //type map is a hash table, which we call dict cause we come from python. We create a slice of them, in order to have 3

	//Compute 1,2,3 order Shannon frequencies of the provided words (including space) and put them in a map
	//We currently have no use for 1-grams, so we could avoid to compute them.
	for n := 0; n < 3; n += 1 {
		// Inizialize the n map and submaps since nil maps cannot be written to. This is a drawback of using maps of maps
		// It's simpler to use a switch than to unravel the combinatorial complexity, and probably performs better
		switch n {
		case 0:
			dict[0] = make(map[string]map[string]int, 1)
			dict[0][""] = make(map[string]int)
		case 1:
			dict[1] = make(map[string]map[string]int, len(dict[0][""]))
			for c := range dict[0][""] {
				dict[1][c] = make(map[string]int)
			}
		case 2:
			dict[2] = make(map[string]map[string]int)
			for c := range dict[0][""] {
				for d := range dict[1][c] {
					dict[2][c+d] = make(map[string]int)
				}
			}
		}
		for i := 0; i < len(arr)-n; i += 1 {
			s = arr[i : i+n+1]
			a = s[0:n]
			b = s[n:]
			dict[n][string(a)][string(b)]++
		}
		//fmt.Println("Dict:", dict[n])
	}

	for n := 0; n < numWords; n++ {
		fmt.Println(generateWord(dict))
	}
}
