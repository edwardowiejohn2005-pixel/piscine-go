package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	// compute deck length without using len()
	total := 0
	for range deck {
		total++
	}

	// assume deck is divisible by 4 as per the task (12 cards → 3 each)
	cardsPerPlayer := total / 4

	for i := 0; i < 4; i++ {
		start := i * cardsPerPlayer
		end := start + cardsPerPlayer

		fmt.Printf("Player %d: ", i+1)
		for j := start; j < end; j++ {
			// print with commas between cards, no trailing comma
			if j != end-1 {
				fmt.Printf("%d, ", deck[j])
			} else {
				fmt.Printf("%d", deck[j])
			}
		}
		fmt.Printf("\n")
	}
}
