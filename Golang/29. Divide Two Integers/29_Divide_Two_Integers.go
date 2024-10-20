
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := (dividend < 0) != (divisor < 0)

	absDividend := int64(math.Abs(float64(dividend)))
	absDivisor := int64(math.Abs(float64(divisor)))

	result := 0

	for absDividend >= absDivisor {
		tempDivisor := absDivisor
		multiple := 1

		for absDividend >= (tempDivisor << 1) {
			tempDivisor <<= 1
			multiple <<= 1
		}

		absDividend -= tempDivisor
		result += multiple
	}

	if negative {
		return -result
	}
	return result
}


