func sortColors(nums []int)  {
    for j:=0; j < len(nums); j++ {
        for i := 0; i < len(nums)-1; i++ {
            if nums[i] > nums[i+1]{  
            nums[i], nums[i+1] = nums[i+1], nums[i]
            }
        }
    } 
}
