func longestPalindrome(s string) string {
    if len(s) < 1 {
        return ""
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        len1 := expandAroundCenter(s, i, i)
        len2 := expandAroundCenter(s, i, i+1)
        length := max(len1, len2)
        if length > end-start {
            start = i - (length-1)/2
            end = i + length/2
        }
    }
    return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
