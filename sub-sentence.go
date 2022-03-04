// Input: the sentence is not a sentence not is an apple not yes a
// N: 5 , Int: > 0

// Output: the sentence not a sentence
// Output: is an apple not a

// Map

// the: [sentence]
// sentence: [is, not]
// is: [not, an]
// not: [a, is, yes]

// O(n) + O(N)

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const N = 5

func main() {
	str := "the sentence is not a sentence not is an apple not yes a"
	split := strings.Split(str, " ")
	fmt.Println(split)
	//fmt.Println("The length of the slice is:", len(split))

	elementMap := make(map[string][]string)
	for i := 0; i < len(split)-1; i++ {
		elementMap[split[i]] = append(elementMap[split[i]], split[i+1])
	}
	lastWordIndex := len(split) - 1
	if _, ok := elementMap[split[lastWordIndex]]; !ok {
		//do something here
		elementMap[split[lastWordIndex]] = []string{}
	}
	//fmt.Println(elementMap)

	outSlice := make([]string, N)

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	curIndex := rand.Intn(len(split))
	curWord := split[curIndex]
	outSlice[0] = curWord
	//fmt.Println("outSlice[0]: ", curWord)

	i := 1
	var nextIndex int
	var nextWord string
	for i < N {
		if len(elementMap[curWord]) == 0 {
			nextIndex = rand.Intn(len(split))
			nextWord = split[nextIndex]
		} else {
			//fmt.Println("[]string{}: ", elementMap[curWord])
			nextIndex = rand.Intn(len(elementMap[curWord]))
			nextWord = elementMap[curWord][nextIndex]
		}
		outSlice[i] = nextWord

		curWord = nextWord
		//	out := fmt.Sprintf("outSlice[%d]: %v", i, curWord)
		//fmt.Println(out)
		i++
	}

	//fmt.Println(outSlice)
	output := strings.Join(outSlice, " ")
	fmt.Println(output)

}

