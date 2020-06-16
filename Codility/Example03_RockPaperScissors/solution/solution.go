package solution

import (
	"fmt"
)

// Solution returns the max. number of points a player scores playing RPS using a specific strategy
/* Franco is going to play a game of Rock-Paper-Scissors with his friend Giovanni. In each turn, both players make their chosen gesture (rock, paper, scissors) simultaneously. After every turn, players gain points as follows: 2 for a win (rock beats scissors, scissors beats paper, paper beats rock), 1 for a tie (when both players show the same gesture) and 0 for a loss.
Franco wants to surprise Giovanni by using a very simple strategy: he will make the same gesture in every turn throughout the game. What is the max. number of points he can score using this strategy?
Write a function solution(G) that, given a string G representing the sequence of Giovanni's turns ('R' represents rock, 'P' paper, 'S' scissors), returns the max. number of points Franco can score using his strategy.
For example:
- Given "RSPRS", the function should return 6 (with the chosen gesture being rock). Franco will gain 6 points (he will win in the 2nd and 5th turns, and tie in the 1st and 4th turns).
- Given "SRR" the function should return 4 (chosen gesture Rock; he will win in the 1st turn and tie in the 2nd and 3rd turns).
- Given "PRPRRPP" the function should return 10 (chosen gesture Paper).
- Give "PPPPRRSSSSS" the function should return 13 (chosen gesture Scissors).

Assume that:
- the length of G is within the range [1...100];
- string G consists only of letters 'R', 'P', 'S'.
*/
func Solution(G string) int {
	b := []byte(G)
	// creating a map "rpsCount" to count the occurences of each of the 'R', 'P', 'S'
	rpsCount := map[rune]int{
		'R': 0,
		'P': 0,
		'S': 0,
	}
	for i, value := range b {
		switch value {
		case 'R':
			rpsCount['R']++
		case 'P':
			rpsCount['P']++
		case 'S':
			rpsCount['S']++
		default:
			fmt.Printf("No valid letter %c in position %d", b[i], i)
		}
	}
	// select the option RPS that will result in the max. number of points
	FrancoOption := keyWithMaxOccurences(rpsCount)
	fmt.Printf("Franco's option is %c\n", FrancoOption)
	return score(rpsCount, FrancoOption)
}

func keyWithMaxOccurences(count map[rune]int) int {
	// finding the key with the max. occurences
	Count := count['R'] //TODO: remove hardcoded key
	Option := 'R'
	for key := range count {
		if Count < count[key] {
			Count = count[key]
			Option = key
		}
	}
	return int(Option)
}

func score(count map[rune]int, option int) int {
	scoreCalc := 0
	for range count {
		switch option {
		case 'R':
			scoreCalc = count['R']*1 + count['S']*2
		case 'P':
			scoreCalc = count['P']*1 + count['R']*2
		case 'S':
			scoreCalc = count['S']*1 + count['P']*2
		default:
			fmt.Printf("Error: score was not calculated. Please try again.")
		}
	}
	return scoreCalc
}
