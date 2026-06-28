func search(nums []int, target int) int {

    start := 0
    end := len(nums)-1
    mid := start + (end - start)/2

    for start <= end {
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            //right side
            start = mid +1
        } else if nums[mid] > target {
            end = mid -1
        }
        mid = start + (end - start)/2
    }
    return -1
}
