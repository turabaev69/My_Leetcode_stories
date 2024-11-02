func longestPalindrome(s string) int {
    charCount := make(map[rune]int)
    for _, ch := range s {
        charCount[ch]++
    }
    length := 0
    oddFound := false
    for _, count := range charCount {
        if count%2 == 0 {
            length += count
        } else {
            length += count - 1
            oddFound = true
        }
    }
    if oddFound {
        length++
    }
    return length
}
