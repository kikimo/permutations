package main

import (
	"fmt"
	"os"
	"strconv"
)

type DXLNode struct {
	prev *DXLNode
	next *DXLNode
	val int
}

type DXList struct {
	head DXLNode
	tail DXLNode
	size int
}

func (o *DXLNode) Hide() {
	prev := o.prev
	next := o.next
	prev.next = next
	next.prev = prev
}

func (o *DXLNode) Show() {
	prev := o.prev
	next := o.next
	prev.next = o
	next.prev = o
}

func (l *DXList) Append(o *DXLNode) {
	o.next = &(l.tail)
	o.prev = l.tail.prev
	o.prev.next = o
	o.next.prev = o
}

func New() *DXList {
	l := &DXList{}
	l.head.prev = nil
	l.head.next = &(l.tail)
	l.tail.next = nil
	l.tail.prev = &(l.head)

	return l
}

func doPerm(l *DXList) [][]int {
	ret := [][]int{}
	for p := l.head.next; p != &(l.tail); p = p.next {
		v := p.val
		p.Hide()
		subRet := doPerm(l)
		if len(subRet) == 0 {
			ret = append(ret, []int{v})
		} else {
			for i := range subRet {
				e := append(subRet[i], v)
				ret = append(ret, e)
			}
		}

		p.Show()
	}

	return ret
}

func Perm(nums []int) [][]int {
	l := New()
	for i := range nums {
		// fmt.Printf("appending: %d\n", nums[i])
		l.Append(&DXLNode{val: nums[i]})
	}

	// for p := l.head.next; p != &(l.tail); p = p.next {
	// 	fmt.Printf("val: %d\n", p.val)
	// }

	return doPerm(l)
}

func main() {
	ns := os.Args[1]
	n, err := strconv.Atoi(ns)
	if err != nil {
		panic(err)
	}

	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, i + 1)
	}
	ret := Perm(nums)
	fmt.Printf("sz: %d\n", len(ret))
}

