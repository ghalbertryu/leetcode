package dp

import (
	"log"
	"testing"
)

func TestCoinChange(t *testing.T) {
	log.Println(coinChange([]int{1, 2, 5}, 100))
	log.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}
