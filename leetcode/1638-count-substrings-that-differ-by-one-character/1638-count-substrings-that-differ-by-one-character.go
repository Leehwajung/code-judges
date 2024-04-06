package main

func countSubstrings(s string, t string) int {
	answer := 0
	for sIdx, sc := range s {
		for tIdx, tc := range t {
			if sc == tc {
				continue
			}
			dscCnt := 1
			si, ti := sIdx, tIdx
			for si > 0 && ti > 0 {
				si--
				ti--
				if s[si] != t[ti] {
					break
				}
				dscCnt++
			}
			ascCnt := 1
			si, ti = sIdx, tIdx
			for si < len(s)-1 && ti < len(t)-1 {
				si++
				ti++
				if s[si] != t[ti] {
					break
				}
				ascCnt++
			}
			answer += dscCnt * ascCnt
		}
	}
	return answer
}
