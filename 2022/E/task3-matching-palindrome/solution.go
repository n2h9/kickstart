package main

import "fmt"

func main() {
	var t byte
	fmt.Scanf("%d", &t)

	var i byte
	for i = 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)

		var s string
		fmt.Scanf("%s\n", &s)

		res := findMatchingPalindrome(s)
		fmt.Printf("Case #%d: %s\n", i, res)
	}
}

func findMatchingPalindrome(s string) string {
	subs := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		subs = append(subs, s[i])
		if isPalindromeBytes(subs) && isPalindrome(s+string(subs)) {
			return string(subs)
		}
	}

	return s
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeBytes(s []byte) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
