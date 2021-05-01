// Go Program to print all the permutations of the given string
package main

import "fmt"

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		//fmt.Println("i val print - ", i)
		f(a) // will just print this is a closure function
		return
	}
	//fmt.Println("i value pre recursion - ", i)
	perm(a, f, i+1)
	//fmt.Println("i value post recursion - ", i)
	for j := i + 1; j < len(a); j++ {
		//fmt.Printf("i - %d j - %d value\n", i,j)
		a[i], a[j] = a[j], a[i]  // swap it to make permutations like abc to acb
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]  // restore it again for the next call
	}
}

func main() {
	Perm([]rune("abc"), func(a []rune) {
		fmt.Println(string(a))
	})
}
