package main

import "fmt"

// You are given two string arrays words1 and words2.

// A string b is a subset of string a if every letter in b occurs in a including multiplicity.

// For example, "wrr" is a subset of "warrior" but is not a subset of "world".
// A string a from words1 is universal if for every string b in words2, b is a subset of a.

// Return an array of all the universal strings in words1. You may return the answer in any order.

// Example 1:

// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
// Output: ["facebook","google","leetcode"]
// Example 2:

// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
// Output: ["apple","google","leetcode"]

// Things to note:
// all words are max 10 char
// list of words are 10,000 max with no duplicates
// Idea 1:
// 1. First We can make a hashmap of word to letter count for w in words2 and store in array. We use hashmap here since we will be repeatedly
//  needting the letter count for each word in thsis array. faster lookup - Time complexity O(n) * m where m is # words in w2 and n is length
// 2. for each word in words1, we also create hashmap and compare key counts to key counts of each word in w2
// This above solution has time limit exceeded. We can simplify this by not having to check
// each word in words1 against each letter count for each word in words2.

// Optimal Solution:
// Lets Just keep track of the max count seen for each letter out of all words in words2. This way we can immediately elimate
// which ever word doesnt apply instead of going through all words in words2 each time

func getCount(word string) []int {
	count := make([]int, 26)
	for _, char := range word {
		count[char-'a'] += 1
	}
	return count
}

func wordSubsets(words1 []string, words2 []string) []string {
	var res []string
	targetCounts := make([]int, 26)
	// we want to get the max count for every letter for each word in words2
	for _, word2 := range words2 {
		w2 := getCount(word2)
		for idx, _ := range targetCounts {
			if targetCounts[idx] < w2[idx] {
				targetCounts[idx] = w2[idx]
			}
		}
	}

	for _, w := range words1 {
		w1 := getCount(w)
		isSub := true
		for idx, _ := range w1 {
			// there are not enough letters in word1 to meet the targe count
			if w1[idx] < targetCounts[idx] {
				isSub = false
				break
			}
		}
		if isSub == true {
			res = append(res, w)
		}
	}

	return res
}

func main() {
	fmt.Printf("%v", wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"le", "lo"}))

}

// My Original code which got TLE:

// func isSubset(m1, m2 map[rune]int) bool {
// 	for k2, v2 := range m2 {
// 		if v1, ok := m1[k2]; ok && v2 <= v1 {
// 			continue
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

// func getCount(word string) map[rune]int {
// 	count := make(map[rune]int)
// 	for _, char := range word {
// 		if _, ok := count[char]; ok {
// 			count[char] += 1
// 		} else {
// 		    count[char] = 1
//         }
// 	}
// 	return count
// }

// func wordSubsets(words1 []string, words2 []string) []string {
// 	var res []string
// 	counts2 := make([]map[rune]int, len(words2))
// 	for idx, w := range words2 {
// 		counts2[idx] = getCount(w)
// 	}

// 	for _, w := range words1 {
// 		c1 := getCount(w)
// 		isSub := true
// 		for _, c2 := range counts2 {
// 			if !isSubset(c1, c2) {
// 				isSub = false
// 			}
// 		}
// 		if isSub == true {
// 			res = append(res, w)
// 		}
// 	}

// 	return res
// }
