package main

import (
	"bufio"
	"fmt"
	"os"
)

type listNode struct {
	str  string
	next *listNode
}

func lenList(h *listNode) int {
	q := h
	l := 0
	for q != nil {
		l++
		q = q.next
	}
	return l
}

func printList(h *listNode, out *bufio.Writer) {
	q := h
	for q != nil {
		if q.next != nil {
			fmt.Fprintf(out, "%s ", q.str)
		} else {
			fmt.Fprintf(out, "%s", q.str)
		}
		q = q.next
	}
	fmt.Fprint(out, "\n")
}

func findInList(s string, h *listNode) *listNode {
	q := h
	for q != nil {
		if q.str == s {
			return q
		}
		q = q.next
	}
	return nil
}

func delNode(s string, h *listNode) *listNode {
	var p *listNode = nil
	q := h
	for q != nil {
		if q.str == s {
			if q == h {
				h = q.next
			} else {
				p.next = q.next
			}

		}
		p = q
		q = q.next
	}
	return h
}

func hash(s string, m int, p int, x int) int {
	f := 0
	power := 1
	for i := 0; i < len(s); i++ {
		f = (f + int(s[i])*power) % p
		power = (power * x) % p
	}
	return f % m
}

func main() {
	inFile, _ := os.Open("input.txt")

	defer inFile.Close()

	in := bufio.NewReader(inFile)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m, n int
	fmt.Fscan(in, &m)
	hashMap := make(map[int]*listNode)
	fmt.Fscan(in, &n)

	var command, s string
	var i, h int
	for j := 0; j < n; j++ {
		fmt.Fscan(in, &command)
		switch command {
		case "add":
			fmt.Fscan(in, &s)
			h = hash(s, m, 1000000007, 263)
			if head, ok := hashMap[h]; !ok {
				hashMap[h] = &listNode{s, nil}
			} else {
				if p := findInList(s, head); p == nil {
					hashMap[h] = &listNode{s, hashMap[h]}
				}
			}
		case "del":
			fmt.Fscan(in, &s)
			h = hash(s, m, 1000000007, 263)
			if head, ok := hashMap[h]; ok {
				if p := findInList(s, head); p != nil {
					if lenList(head) == 1 {
						delete(hashMap, h)
					} else {
						hashMap[h] = delNode(s, head)
					}
				}
			}
		case "find":
			fmt.Fscan(in, &s)
			h = hash(s, m, 1000000007, 263)
			if head, ok := hashMap[h]; ok {
				if p := findInList(s, head); p != nil {
					fmt.Fprintln(out, "yes")
				} else {
					fmt.Fprintln(out, "no")
				}
			} else {
				fmt.Fprintln(out, "no")
			}

		case "check":
			fmt.Fscan(in, &i)
			if head, ok := hashMap[i]; ok {
				printList(head, out)
			} else {
				fmt.Fprint(out, "\n")
			}
		}
	}
}
