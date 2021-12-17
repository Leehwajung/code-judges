package main

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	resultTable := make([]int, amount+1)
	for i := 1; i < len(resultTable); i++ {
		resultTable[i] = -1
	}

	for _, coin := range coins {
		for currAmount := coin; currAmount <= amount; currAmount++ {
			subResult := resultTable[currAmount-coin]
			if subResult < 0 {
				continue
			}
			result := resultTable[currAmount]
			if result < 0 || subResult+1 < result {
				resultTable[currAmount] = subResult + 1
			}
		}
	}
	return resultTable[amount]
}
