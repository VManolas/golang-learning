package solution

import "fmt"

// Solution returns the day of the week that is K days later than the S day
// Days of the week are represented as three-letter strings ("Mon", "Tue", ...)
// Write a function solution() that, given a string S representing the day of the week and an integer K (between 0 and 500, inclusive) returns the day of the week that is K days later.
// For example, given S = "Wed" and K = 2, the function should return "Fri"
// Given S = "Sat" and K = 23, the function should return "Fri"
func Solution(days []string, S string, K int) string {
	// boolean to validate wheter S is a valid option (it it expected to be)
	validDay := false

	for i, value := range days {
		fmt.Println(value, S)
		if S == value {
			validDay = true
			K += i
			break
		}
	}

	if validDay == false {
		return "Error: the day received (S string) is not represented in a valid 3-letter way"
	}

	fmt.Println(K)
	for K >= 7 {
		K = K % 7
	}
	return days[K]
}
