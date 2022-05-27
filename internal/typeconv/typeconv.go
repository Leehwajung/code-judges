package typeconv

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fatalExit(err)
		return 0
	}
	return i
}

func BitStringToUint32(s string) uint32 {
	i, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		fatalExit(err)
		return 0
	}
	return uint32(i)
}

func StringsToInts(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		i := StringToInt(s)
		ints = append(ints, i)
	}
	return ints
}

func StringsToIntPairs(strs []string) [][]int {
	if len(strs)%2 != 0 {
		fatalExit(errors.New("invalid integer pair as arguments passed"))
		return nil
	}
	intPairs := make([][]int, 0, len(strs)/2)
	for idx := 0; idx < len(strs); idx += 2 {
		i1 := StringToInt(strs[idx])
		i2 := StringToInt(strs[idx+1])
		intPairs = append(intPairs, []int{i1, i2})
	}
	return intPairs
}

func StringsToIntGrid(strs []string, rowLen int) [][]int {
	if len(strs)%rowLen != 0 {
		fatalExit(errors.New("invalid integer grid as arguments passed"))
		return nil
	}

	intGrid := make([][]int, 0, len(strs)/rowLen)
	intRow := []int(nil)
	for idx := 0; idx < len(strs); idx++ {
		if idx%rowLen == 0 {
			intRow = make([]int, 0, rowLen)
		}
		i := StringToInt(strs[idx])
		intRow = append(intRow, i)
		if idx%rowLen == rowLen-1 {
			intGrid = append(intGrid, intRow)
		}
	}
	return intGrid
}

func StringsToSet(strs []string) map[string]bool {
	set := make(map[string]bool, len(strs))
	for _, s := range strs {
		set[s] = true
	}
	return set
}

func fatalExit(err error) {
	_, _ = fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
