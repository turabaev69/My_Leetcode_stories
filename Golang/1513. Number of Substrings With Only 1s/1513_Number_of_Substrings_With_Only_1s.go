package main

func numSub(s string) int {
    const mod = 1_000_000_007
    count, result := 0, 0

    for _, ch := range s {
        if ch == '1' {
            count++
            result = (result + count) % mod
        } else {
            count = 0
        }
    }

    return result
}
