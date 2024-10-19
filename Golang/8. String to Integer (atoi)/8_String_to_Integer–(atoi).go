import "math"

func myAtoi(s string) int {
    i, sign, result := 0, 1, 0
    n := len(s)
    for i < n && s[i] == ' ' {
        i++
    }
    if i < n && (s[i] == '+' || s[i] == '-') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }
    for i < n && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')
        if result > (math.MaxInt32-digit)/10 {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }
        result = result*10 + digit
        i++
    }
    return result * sign
}
