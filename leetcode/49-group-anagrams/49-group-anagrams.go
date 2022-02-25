package main

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string)
	for _, str := range strs {
		key := [26]byte{}
		for _, c := range str {
			key[c-'a']++
		}
		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, anagrams := range groups {
		result = append(result, anagrams)
	}
	return result
}
