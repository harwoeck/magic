package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/harwoeck/magic"
)

type tx struct {
	id     string
	from   []txin
	to     [1]txout
	amount float64
}

func (t tx) Print() {
	fmt.Printf("tx[%s] with %d ins and %d outs has %f cashflow\n    ins\n", t.id, len(t.from), len(t.to), t.amount)
	for _, tin := range t.from {
		fmt.Printf("    txin[%s] with submitTime %d\n", tin.id, tin.submitTime)
	}
	fmt.Printf("    outs\n")
	for _, tout := range t.to {
		fmt.Printf("    txout[%s] with amount %d and xs (l1=[%v %v], l2=[%v %v])\n", tout.owner, tout.amount, tout.xs[0][0], tout.xs[0][1], tout.xs[1][0], tout.xs[1][1])
	}
}

type txin struct {
	id         string
	submitTime int64
}

type txout struct {
	owner  string
	amount int64
	child  *txin
	xs     [2][2]bool
}

func main() {
	src := bufio.NewScanner(strings.NewReader(
		"4\n" +
			"0x300 2 0x100 150 0x200 250 Alfonso 400 0x000 1 true true false true 50\n" +
			"0x600 1 0x100 400 Stephan 100 0x000 1 true false false true 33.7\n" +
			"str1 2 str2 100 str3 100 Name 200 xx 1 true true true true 200\n" +
			"str4 1 str5 100 OtherName 20 xx 1 false false false false 100"))
	m := magic.NewManager(src)

	for _, t := range m.Read([]*tx{}).([]*tx) {
		t.Print()
	}
}
