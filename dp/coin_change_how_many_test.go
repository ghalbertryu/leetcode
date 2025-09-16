package dp

import (
	"log"
	"testing"
)

func TestCoinChangeHowManyWays(t *testing.T) {
	log.Println(coinChangeHowManyWays([]int{1, 2, 5}, 1)) // 1
	log.Println(coinChangeHowManyWays([]int{1, 2, 5}, 2)) // 2
	log.Println(coinChangeHowManyWays([]int{1, 2}, 4))    // 3
}
