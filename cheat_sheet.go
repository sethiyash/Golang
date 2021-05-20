

package main

import (
	"fmt"
	"math"
)

// Use Math package to get the max int and min int we have MaxInt64 also 
const maxInt = math.MaxInt32  
const minInt = math.MinInt32

//If you want to pass a slice as a parameter to a function, and have that function modify the original slice, 
//then you have to pass a pointer to the slice:

func myAppend(list *[]string, value string) {
    *list = append(*list, value)
}




