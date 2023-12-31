// # Alice has n balloons arranged on a rope. You are given a 0-indexed string colors where colors[i] is the color of the ith balloon.

// # Alice wants the rope to be colorful. She does not want two consecutive balloons to be of the same color, so she asks Bob for help. Bob can remove some balloons from the rope to make it colorful. You are given a 0-indexed integer array neededTime where neededTime[i] is the time (in seconds) that Bob needs to remove the ith balloon from the rope.

// # Return the minimum time Bob needs to make the rope colorful.
package main

import "math"

func minCost(colors string, neededTime []int) int {
	if len(colors) < 2 {
		return 0
	}

	currIdx := 1
	var total float64
	for currIdx < len(colors) {
		// case 1 - we find a repeated color
		if colors[currIdx] == colors[currIdx-1] {
			// we want to keep the color with the max neededTime. So lets just
			// keep track of the sum of neededTime of all colors then subtract the max at the end
			p := currIdx
			maxTime := float64(neededTime[currIdx-1])
			colorTotal := float64(neededTime[currIdx-1])
			for p < len(colors) && colors[p] == colors[p-1] {
				maxTime = math.Max(float64(maxTime), float64(neededTime[p]))
				colorTotal += float64(neededTime[p])
				p += 1
			}
			total += colorTotal - maxTime
			currIdx = p + 1
			continue
		}
		// case 2 - no repeated color
		currIdx += 1
	}

	return int(total)
}
func main() {
}
