// declear and init er jonno :=
// otherwise = to init only

package main

func MaxSubArraySum(arr []int) int {
	arrLen := len(arr)
	currentSum := arr[0]
	maxSum := arr[0]

	for i := 1; i < arrLen; i++ {
		currentSum = max(arr[i], currentSum+arr[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// func main1(){
// 	testArr := []int{3,2,4,5}
// 	fmt.Println(maxSubArraySum(testArr))
// }
