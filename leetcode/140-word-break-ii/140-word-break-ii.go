package main

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	dict := make(map[byte][]string, 26)
	for _, word := range wordDict {
		dict[word[0]] = append(dict[word[0]], word)
	}

	sentences := makeSentences(s, dict)

	results := make([]string, len(sentences))
	for i, sentence := range sentences {
		results[i] = strings.Join(sentence, " ")
	}
	return results
}

func makeSentences(s string, dict map[byte][]string) (sentences [][]string) {
	if len(s) == 0 {
		return [][]string{nil}
	}

	for _, word := range dict[s[0]] {
		if len(word) > len(s) {
			continue
		} else if s[:len(word)] != word {
			continue
		}

		tails := makeSentences(s[len(word):], dict)
		for _, tail := range tails {
			sentence := append([]string{word}, tail...)
			sentences = append(sentences, sentence)
		}
	}
	return sentences
}
