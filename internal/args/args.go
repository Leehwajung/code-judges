package args

import (
	"errors"
	"fmt"
	"os"
)

func MustEnough(minCount int) {
	if len(os.Args)-1 < minCount {
		fatalExit(errors.New("not enough arguments passed"))
	}
}

func MustEvenCount() {
	if (len(os.Args)-1)%2 != 0 {
		fatalExit(errors.New("odd arguments passed"))
	}
}

func fatalExit(err error) {
	_, _ = fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
