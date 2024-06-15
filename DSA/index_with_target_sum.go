package main

import "fmt"
// Output the pair of index of the array elements which when added together will result in the given target sum
// t{1,2,3,4,5,6,7,8,9};  target=9  // observe that the sorted array is given

// we will be using two pointer approach
// this approach only work when the array is sorted

func main()  {
	arr := []int{1,1,2,3,4,5,6,7,8,9}
	target := 9
	rightP := len(arr)-1

	for leftP:=0; leftP < len(arr) && rightP > leftP; {
		if arr[leftP]+arr[rightP] == target {
			fmt.Printf("%d, %d\n",leftP, rightP) // Printing index for which the sum is 9
			leftP++
		} else if arr[leftP]+arr[rightP] > target {
			rightP--
		} else if arr[leftP]+arr[rightP] < target {
			leftP++
		}
	}
	return
}
