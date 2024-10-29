func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if index, found := hashMap[complement]; found {
            return []int{index, i}
        }
        hashMap[num] = i
    }
    return nil 
}
