package main

import (
	"bufio"
	"fmt"
	"os"
)

/*func hash(s string, m int, p int, x int) int {
	f := 0
	power := 1
	for i := 0; i < len(s); i++ {
		f = (f + int(s[i])*power) % p
		power = (power * x) % p
	}
	return f % m
}*/

func hash(s string) int {
	f := 0
	for i := 0; i < len(s); i++ {
		f += int(s[i])
	}
	return f
}

func check(a string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var pattern, text string

	fmt.Fscan(in, &pattern)
	fmt.Fscan(in, &text)

	textLen := len(text)
	patternLen := len(pattern)

	patternHash := hash(pattern)
	textHash := hash(text[:patternLen])

	if textHash == patternHash {
		if check(pattern, text[:patternLen]) {
			fmt.Fprintf(out, "%d ", 0)
		}
	}

	for i := 1; i <= textLen-patternLen; i++ {
		//fmt.Fprintln(out, text[i:i+patternLen])
		textHash = textHash - int(text[i-1]) + int(text[i+patternLen-1])
		if textHash == patternHash {
			if check(pattern, text[i:i+patternLen]) {
				fmt.Fprintf(out, "%d ", i)
			}
		}
	}

}
