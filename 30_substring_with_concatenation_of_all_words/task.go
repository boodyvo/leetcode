package task

import "maps"

func findSubstring(s string, words []string) []int {
	wordsNum := len(words)
	wordLen := len(words[0])
	strLen := len(s)

	response := make([]int, 0)
	templateUsed := make(map[string]int, wordsNum)
	for _, word := range words {
		templateUsed[word]++
	}

	for i := 0; i <= strLen-wordLen*wordsNum; i++ {
		rejected := false
		used := maps.Clone(templateUsed)

		for j := i; j < i+wordLen*wordsNum; j += wordLen {
			used[s[j:j+wordLen]]--
			if used[s[j:j+wordLen]] < 0 {
				rejected = true

				break
			}
		}

		if !rejected {
			response = append(response, i)
		}
	}

	return response
}
