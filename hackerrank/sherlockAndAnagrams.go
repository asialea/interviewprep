package main

import (
	"fmt"
	"sort"
)

// Two strings are anagrams of each other if the letters of one string can be rearranged to form the other string.
// Given a string, find the number of pairs of substrings of the string that are anagrams of each other.
// Example

// The list of all anagrammatic pairs is  at positions  respectively.
// Function Description
// Complete the function sherlockAndAnagrams in the editor below.
// sherlockAndAnagrams has the following parameter(s):
// string s: a string
// Returns
// int: the number of unordered anagrammatic pairs of substrings in

// ideas
// 1. go through substrings, keep track of word count for each substring in hashmap, this would involve checking if one
// hashmap is equal to another, can we even do that since they are unordered sets... lets try
// sort each substring and see if we have a matching substring in the map
// next, we are counting combinations of 2, so for each match we find we need to see how many groups of 2 we can make
// time complexity = n * n * nlogn = n^2 where n is number of chars in string
// space complexity =

func sherlockAndAnagrams(str string) int32 {
	var total int32
	windowSize := 1

	// func to sort string
	sorted := func(word string) string {
		s := []rune(word)
		sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
		return string(s)
	}

	// get all substrings of each window size
	for windowSize < len(str) {
		seen := make(map[string]int32)
		for x := 0; x+windowSize <= len(str); x++ {
			unsortedStr := str[x : x+windowSize]
			sortedStr := sorted(unsortedStr)
			if _, ok := seen[sortedStr]; ok {
				seen[sortedStr] += 1
			} else {
				seen[sortedStr] = 1
			}
		}
		// count number of combinations we can make using formula n(n-1)/2
		for _, val := range seen {
			if val > 0 {
				total += (val * (val - 1)) / 2
			}
		}
		windowSize++
	}
	return total
}

func main() {
	fmt.Println(sherlockAndAnagrams("kkkk"))
}
