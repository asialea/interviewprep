package main
import "fmt"
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
// Example 1:
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
//
// Example 2:
// Input: n = 1
// Output: ["()"]

// so at each point we have the following cases...
// numOpen numClose
// if leftOpen == leftClose == 0: return string
// add an open .. if leftOpen > 0
// add a close .. if leftOpen < leftClose ... aka we  have one or more open brackets placed

func helper(leftOpen, leftClose int, result *[]string, currString string) {
	// base case
	if leftOpen == 0 && leftClose == 0 {
		// When you pass a pointer to function, the function gets a copy of it, 
		// So assigning new value to the given pointer will not result in assigning it to the original one.
		// therefore we use *result to get the memory address of the orignal pointer and assign a value to it
		*result = append(*result, currString) 
		return
	}
	
	// Add an open next
	if leftOpen > 0 {
		helper(leftOpen-1, leftClose, result, currString + "(")
	}

	// Add a closed next
	if leftOpen < leftClose {
		helper(leftOpen, leftClose-1, result, currString + ")")
	}

}
 
func generateParentheses(n int) []string{
	var result []string
	helper(n, n, &result, "")
	return result
}

func main() {
	fmt.Println(generateParentheses(3))
}