// find longest substring with K unique characters in given string
// str = aabbcc k = 1
// Output - aa/bb/cc (print any longest substring with K unique characters)
// Run on - https://go.dev/play/p/NyDUusiir8b

package main

import "fmt"

func longestSubstring(s string, k int) string {

	mp := make(map[byte]int)
	i := 0
	ans := -1
	left := 0
	right := 0
	res := ""
	for j := 0; j < len(s); j++ {
		mp[s[j]] = mp[s[j]] + 1
		for len(mp) > k {
			mp[s[i]]--
			if mp[s[i]] == 0 {
				delete(mp, s[i])
			}
			i++
		}
		if len(mp) == k {
			if j-i+1 > ans {
				left = i
				right = j
			}
			ans = max(ans, j-i+1)
		}
	}
	for itr := left; itr <= right; itr++ {
		res += string(s[itr])
	}
	return res
}

func main() {
	s := "aabacbebebebe"
	k := 3
	fmt.Println(longestSubstring(s, k)) // output - cbebebebe
}
