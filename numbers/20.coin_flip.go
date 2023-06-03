package numbers

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Coin Flip Simulation :
// Write some code that simulates flipping a single coin however many times the user decides.
// The code should record the outcomes and count the number of tails and heads.
func CoinFlip() {
	var (
		flipCoinAnswer string
		heads          uint
		tails          uint
	)

	for {
		fmt.Printf("\033[33mDo you want to flip the coin?(y/n)\033[0m ")
		fmt.Scanln(&flipCoinAnswer)
		flipCoinAnswer = strings.ToLower(flipCoinAnswer)

		if flipCoinAnswer == "y" || flipCoinAnswer == "yes" {
			answer := headOrTail()
			fmt.Printf("Coin flipped: \033[32m%s\n\033[0m", answer)
			if answer == "head" {
				heads++
			} else {
				tails++
			}

		} else {
			fmt.Printf("\033[36mHeads\033[0m: \033[32m%d\033[0m \n", heads)
			fmt.Printf("\033[36mTails\033[0m: \033[32m%d\033[0m \n", tails)
			break
		}

	}
}

func headOrTail() string {
	var (
		head = 0
		tail = 1
	)

	rand.Seed(time.Now().Unix())
	number := rand.Intn(2)

	switch number {
	case head:
		return "head"
	case tail:
		return "tail"
	default:
		panic("WTF!")
	}
}
