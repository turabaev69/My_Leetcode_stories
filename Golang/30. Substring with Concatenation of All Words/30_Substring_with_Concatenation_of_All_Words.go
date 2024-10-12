func findSubstring(s string, words []string) []int {
    result := []int{}
    if len(s) == 0 || len(words) == 0 {
        return result
    }

    wordLength := len(words[0])
    wordsCount := len(words)


    wordMap := map[string]int{}
    for _, word := range words {
        wordMap[word]++
    }

    for i := 0; i < wordLength; i++ {
        left := i
        seen := map[string]int{}
        count := 0

        for j := i; j <= len(s)-wordLength; j += wordLength {
            word := s[j : j+wordLength]

            if _, ok := wordMap[word]; ok {
                seen[word]++
                count++

                for seen[word] > wordMap[word] {
                    leftWord := s[left : left+wordLength]
                    seen[leftWord]--
                    count--
                    left += wordLength
                }

                if count == wordsCount {
                    result = append(result, left)
                }
            } else {
                seen = map[string]int{}
                count = 0
                left = j + wordLength
            }
        }
    }
    return result
}
