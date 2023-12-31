// HackerLand National Bank has a simple policy for warning clients about possible fraudulent account activity. 
// If the amount spent by a client on a particular day is greater than or equal to  the client's median spending for a trailing number of days, 
// they send the client a notification about potential fraud. 
// The bank doesn't send the client any notifications until they have at least that trailing number of prior days' transaction data.
// Given the number of trailing days  and a client's total daily expenditures for a period of  days, determine the number of times the client will receive a notification over all  days.
// https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting

// median can be found by sorting the input and get the element in the center of the array
// for each slidng window sort and find median

func getMedian(arr []int) int {
	if len(arr < 1) {
		return nil
	}
	
	if len(arr) % 2 == 1 {
		return len(arr=)
	}

}

func activityNotifications(expenditure []int32, d int32) int32 {
    // Write your code here

}