package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	numbers := make(map[int]string)
	var n int
	fmt.Fscan(in, &n)
	var command string
	var num int
	var name string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &command)
		fmt.Fscan(in, &num)
		switch command {
		case "add":
			fmt.Fscan(in, &name)
			numbers[num] = name
		case "find":
			if name, ok := numbers[num]; !ok {
				fmt.Fprintln(out, "not found")
			} else {
				fmt.Fprintln(out, name)
			}
		case "del":
			if _, ok := numbers[num]; ok {
				delete(numbers, num)
			}

		}
	}

}
