package problems_test

import (
	"testing"

	pb "github.com/georgesafta/algorithms/problems"
)

func TestBuyStock(t *testing.T) {
	assert(t, []int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97}, 43)
	assert(t, []int{7, 1, 5, 3, 6, 4}, 5)
	assert(t, []int{7, 6, 4, 3, 1}, 0)
	assert(t, []int{100}, 0)
}

func assert(t *testing.T, prices []int, expectedProfit int) {
	profit := pb.BuyStockBruteForce(prices)
	if profit != expectedProfit {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedProfit, profit)
	}

	profit = pb.BuyStockDivideAndConquer(prices)
	if profit != expectedProfit {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedProfit, profit)
	}

	profit = pb.BuyStockDP(prices)
	if profit != expectedProfit {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedProfit, profit)
	}

	profit = pb.BuyStockDPOptimizedSpace(prices)
	if profit != expectedProfit {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedProfit, profit)
	}
}
