package main

// Alice and Bob are opponents in an archery competition. The competition has set the following rules:

// Alice first shoots numArrows arrows and then Bob shoots numArrows arrows.
// The points are then calculated as follows:
// The target has integer scoring sections ranging from 0 to 11 inclusive.
// For each section of the target with score k (in between 0 to 11), say Alice and Bob have shot ak and bk arrows on that section respectively. If ak >= bk, then Alice takes k points. If ak < bk, then Bob takes k points.
// However, if ak == bk == 0, then nobody takes k points.
// For example, if Alice and Bob both shot 2 arrows on the section with score 11, then Alice takes 11 points. On the other hand, if Alice shot 0 arrows on the section with score 11 and Bob shot 2 arrows on that same section, then Bob takes 11 points.

// You are given the integer numArrows and an integer array aliceArrows of size 12, which represents the number of arrows Alice shot on each scoring section from 0 to 11. Now, Bob wants to maximize the total number of points he can obtain.

// Return the array bobArrows which represents the number of arrows Bob shot on each scoring section from 0 to 11. The sum of the values in bobArrows should equal numArrows.

// If there are multiple ways for Bob to earn the maximum total points, return any one of them.

// Input: numArrows = 9,
// aliceArrows = [1,1,0,1,0,0,2,1,0,1,2,0]
// Output:       [0,0,0,0,1,1,0,0,1,2,3,1]

// Input: numArrows = 3,
// aliceArrows = [0,0,1,0,0,0,0,0,0,0,0,2]
// Output: [0,0,0,0,0,0,0,0,1,1,1,0]

// Given number of aarows N, we want to distribute these arrows in such a way that bob scores the max points he can get.
//  but we also want to beat alice i think? this question is unclear

// Idea 1
// - so we want to intelligently distribute the arrows. in any scenario where bob wins, his score only needs to be 1 more than
// the total points Alicse scored for that round
// - can we start from the right side of the arr and just do 1 more than each point that alice has? no this wont work in every case

// Idea 2
// Can we just try every combonation and see which gets us the highest total points and return it? based on if we were too allot
// x arrows at each space?
// This is a backtracking problem . At each point we want to see the cost of winning the round vs losing the round and taking no points.

import "fmt"

func helper(numArrows int, aliceArrows []int, myArrows []int, currIdx int, res *[]int) {
	// base case: we have reached the end of this currArr. i what we observe is a max sum, lets update the result array.
	if currIdx == 12 {
		maxSum := getSum(aliceArrows, *res)
		currSum := getSum(aliceArrows, myArrows)
		if currSum > maxSum {

			points := make([]int, len(myArrows)) // we want to make a copy of current points because otherwise MyArrows will keep changing as we recurse
			copy(points, myArrows)
			// store my arrows as the new result
			*res = points
		}
		return
	}
	
	// case 1 - what would happen if we lost this round, keep our score to 0 for this round
	helper(numArrows, aliceArrows, myArrows, currIdx+1, res)


	// case 2 - we have enough arrows. what would happen if we win this round
	if numArrows >= aliceArrows[currIdx] {
		myArrows[currIdx] = aliceArrows[currIdx] + 1
		helper(numArrows-aliceArrows[currIdx]-1, aliceArrows, myArrows, currIdx+1, res)
		myArrows[currIdx] = 0 
		// ok so here we have to set it back to 0, because we have to back out of this branch and 
		// try the rest of the recursive cases at the previous indexes. 
		// So for example we have
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 1, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 2, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 3, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 4, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 5, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 6, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 7, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 8, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 9, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,alice+1,0], 10, res)
			// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 10, res)
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 11, res)
			//  win_case-> helper(9 - alice+1, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0, alice+1], 11, res))
				//  win_case-> helper(9 - alice+1, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0, alice+1], 12, res)) -> end case
				// back out and go to idx..10 the next point in the recursion
		// helper(9, aliceArrows, [0,0,0,0,0,0,0,0,0,0,0,0], 12, res) -> end case
	}
}

func getSum(aliceArrows, myArrows []int) int {
	total := 0
	for idx, _ := range myArrows {
		if myArrows[idx] > aliceArrows[idx] {
			total += idx
		}
	}
	return total
}

func getCombinations(numArrows int, aliceArrows []int) []int {
	var res []int
	myArrows := make([]int, 12)
	helper(numArrows, aliceArrows, myArrows, 0, &res)
	return res
}

func main() {

	fmt.Println(getCombinations(9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0}))

}
