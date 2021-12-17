package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func StringsToInts(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, s := range strs {
		i := StringToInt(s)
		ints = append(ints, i)
	}
	return ints
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fatalExit(err)
		return 0
	}
	return i
}

func CheckArgsMinCount(count int) {
	if len(os.Args)-1 < count {
		fatalExit(errors.New("not enough arguments passed"))
	}
}

func fatalExit(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
