package main

import "fmt"

// You are given an array and you need to find number of tripets of indices such that the elements at those indices are
// in geometric progression for a given common ratio  and .
// Function Description
// Complete the countTriplets function in the editor below.
// countTriplets has the following parameter(s):
// int arr[n]: an array of integers
// int r: the common ratio
// Returns
// int: the number of triplets
// Input Format
// The first line contains two space-separated integers  and , the size of  and the common ratio.
// The next line contains  space-seperated integers

// idea:
// place all elements in hashmap key: number, value: index  space and time: o(n)
// for each element in original array, check hashmap for the next 2 nums in progression.
// check if index is greater than curr index  time o(n)
// ^ this wont work becasue what if we have multiple geomeetric progressions with overlapping indices ex. [1, 4, 4, 16, 16, 64]
// idea 2:
// lets input these values in a tree. for each node we keep a set of indices where we saw the letter
// then lets do a dfs on the tree ending at depth 3. for each child node make sure the prev node has a greater index then parent.
// if not stop
// Result - i failed here due to time limit exceeded.

func count(level, currIdx, ratio int, seen map[int][]int, currVal int) int32 {
	if level == 3 {
		return 1
	}
	target := ratio * currVal
	indices, ok := seen[target]
	if !ok {
		return 0
	}
	total := int32(0)
	for _, idx := range indices {
		if idx > currIdx {
			total += int32(count(level+1, idx, ratio, seen, target))
		}
	}
	return total
}

func countTriplets(arr []int, ratio int) int32 {
	// First lets index the items in the arr
	seen := make(map[int][]int)
	for idx, val := range arr {
		if _, ok := seen[val]; ok {
			seen[val] = append(seen[val], idx)
		} else {
			seen[val] = []int{idx}
		}

	}
	var total int32
	for idx, el := range arr {
		total += count(1, idx, ratio, seen, el)
	}
	return int32(total)
}

func main() {
	fmt.Println(countTriplets([]int{1, 4, 16, 64}, 4))
}

// REAL SOLUTION PYTHON - use memoization to do it in one pass
// for el on arr
// if el in r3, we have completed a triplet. add value to total
// if el in r2, we know that if we end up finding r*el later in the sequence, we can make r2[v] number ofsequences
// add el*r to r2. if we end up finding this val later in the sequence it can be 1 potential combination

// def countTriplets(arr, r):
//     r2 = Counter()
//     r3 = Counter()
//     count = 0

//     for v in arr:
//         if v in r3:
//             count += r3[v]

//         if v in r2:
//             r3[v*r] += r2[v]

//         r2[v*r] += 1

//     return count
