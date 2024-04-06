package main

func minRemoveToMakeValid(s string) string {
	type pair struct {
		index int
		left  bool
	}
	var invalidity []pair
	for idx, c := range s {
		switch c {
		case '(':
			invalidity = append(invalidity, pair{index: idx, left: true})
		case ')':
			if len(invalidity) > 0 && invalidity[len(invalidity)-1].left {
				invalidity = invalidity[:len(invalidity)-1]
			} else {
				invalidity = append(invalidity, pair{index: idx})
			}
		default:
		}
	}

	if len(invalidity) == 0 {
		return s
	}

	invIdx := 0
	answer := make([]byte, 0, len(s)-len(invalidity))
	for idx, c := range s {
		if invIdx < len(invalidity) && idx == invalidity[invIdx].index {
			invIdx++
			continue
		}
		answer = append(answer, byte(c))
	}
	return string(answer)
}
