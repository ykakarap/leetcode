package main

import (
	"fmt"
)

func main() {

	s1 := "phqghumeaylnlfdxfircvscxggbwkfnqduxwfnfozvsrtkjprepggxrpnrvystmwcysyycqpevikeffmznimkkasvwsrenzkycxf"
	n1 := 1000000
	s2 := "xtlsgypsfadpooefxzbcoejuvpvaboygpoeylfpbnpljvrvipyamyehwqnqrqpmxujjloovaowuxwhmsncbxcoksfzkvatxdknly"
	n2 := 100

	// s1 := "aaa"
	// n1 := 3
	// s2 := "aa"
	// n2 := 1

	fmt.Printf("Answer is : %v \n", getMaxRepetitions3(s1, n1, s2, n2))
}

type WS struct {
	S     string
	N     int
	count int
	pos   int
	i     int
	l     int
}

func (ws *WS) start() {
	ws.l = len(ws.S)
}

func (ws *WS) next() byte {
	r := ws.S[ws.pos]
	if ws.pos == ws.l-1 {
		ws.count++
	}
	ws.i++
	ws.pos = ws.i % ws.l
	return r
}

func (ws *WS) done() bool {
	return (ws.i - 1) == ws.N*ws.l
}

func (ws *WS) reset() {
	ws.i = 0
	ws.count = 0
	ws.pos = 0
}

func (ws *WS) getCount() int {
	return ws.count
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	S1 := &WS{
		S: s1,
		N: n1,
	}
	S1.start()
	S2 := &WS{
		S: s2,
		N: n2,
	}
	S2.start()

	req := S2.next()
	count := 0
	for c := S1.next(); !S1.done(); c = S1.next() {
		if c == req {
			req = S2.next()
			count++
		}
	}

	return (count / S2.l) / S2.N
}

func getMaxRepetitions2(s1 string, n1 int, s2 string, n2 int) int {
	pos1 := 0
	l1 := len(s1)
	count1 := 0

	pos2 := 0
	l2 := len(s2)
	count2 := 0

	for count1 < l1*n1 {
		c := s1[pos1]
		if c == s2[pos2] {

			count2++
			pos2 = count2 % l2
		}
		count1++
		pos1 = count1 % l1
		if count1%(n1*l1) == 0 && count2%(n2*l2) == 0 {
			break
		}
	}

	// fmt.Println(count2)
	// fmt.Println(l2)
	// fmt.Println(n2)
	// fmt.Println(count1)
	// fmt.Println(l1 * n1)

	return (((count2 / l2) / n2) * (l1 * n1)) / count1
}

func getMaxRepetitions3(s1 string, n1 int, s2 string, n2 int) int {
	pos1 := 0
	l1 := len(s1)
	count1 := 0

	pos2 := 0
	l2 := len(s2)
	count2 := 0

	for count1 < l1*n1 {
		c := s1[pos1]
		if c == s2[pos2] {
			count2++
			pos2 = count2 % l2
		}
		count1++
		pos1 = count1 % l1
	}

	return (count2 / l2) / n2
}
