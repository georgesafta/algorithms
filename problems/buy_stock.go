package problems

/**
Find the maximum profit that can be achieved by buying stock once and selling it at a later date.
If not profit can be made, don't buy.
The input for this problem is a list representing the daily price of stock.
*/

// BuyStockBruteForce gives the maximum profit that can be achieved by buying stock once and selling it at a later date.
// Time complexity: O(n^2)
func BuyStockBruteForce(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > max {
				max = profit
			}
		}
	}

	return max
}

// BuyStockDivideAndConquer gives the maximum profit that can be achieved by buying stock once and selling it at a later date.
// Time complexity: O(nlogn)
func BuyStockDivideAndConquer(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	changes := make([]int, len(prices)-1)
	for i := 1; i < len(prices); i++ {
		changes[i-1] = prices[i] - prices[i-1]
	}

	_, _, profit := maxProfit(changes, 0, len(changes)-1)

	return max(profit, 0)
}

// BuyStockDP gives the maximum profit that can be achieved by buying stock once and selling it at a later date.
// Time complexity: O(n)
func BuyStockDP(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	leftMin := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			leftMin[i] = prices[i]
		} else {
			leftMin[i] = min(leftMin[i-1], prices[i])
		}
	}

	rightMax := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		if i == len(prices)-1 {
			rightMax[i] = prices[i]
		} else {
			rightMax[i] = max(rightMax[i+1], prices[i])
		}
	}

	profit := 0
	for i := 0; i < len(prices); i++ {
		profit = max(profit, rightMax[i]-leftMin[i])
	}

	return profit
}

// BuyStockDPOptimizedSpace gives the maximum profit that can be achieved by buying stock once and selling it at a later date.
// Time complexity: O(n)
func BuyStockDPOptimizedSpace(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buy := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		profit = max(profit, prices[i]-buy)
		buy = min(buy, prices[i])
	}

	return profit
}

func maxProfit(changes []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, changes[low]
	}

	mid := low + (high-low)/2
	leftStart, leftEnd, leftProfit := maxProfit(changes, low, mid)
	rightStart, rightEnd, rightProfit := maxProfit(changes, mid+1, high)
	crossStart, crossEnd, crossProfit := maxSumCrossingMid(changes, low, mid, high)

	if leftProfit >= rightProfit && leftProfit >= crossProfit {
		return leftStart, leftEnd, leftProfit
	}

	if rightProfit >= leftProfit && rightProfit >= crossProfit {
		return rightStart, rightEnd, rightProfit
	}

	return crossStart, crossEnd, crossProfit
}

func maxSumCrossingMid(arr []int, left, mid, right int) (int, int, int) {
	sum := arr[mid]
	leftSum := arr[mid]
	leftIndex := mid
	for i := mid - 1; i >= left; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
			leftIndex = i
		}
	}

	sum = arr[mid+1]
	rightSum := arr[mid+1]
	rightIndex := mid + 1
	for i := mid + 2; i <= right; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightSum = sum
			rightIndex = i
		}
	}

	return leftIndex, rightIndex, leftSum + rightSum
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
