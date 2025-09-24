package leetcode

import (
	"log"
	"testing"
)

func TestSuggestedProducts(t *testing.T) {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	log.Println(suggestedProducts(products, "mouse"))
}
