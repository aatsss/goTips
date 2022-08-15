// 双方向リスト, 挿入, 削除, 照会 のジェネリクス

/*
型を any から comparable に変えれば多分比較演算使える
*/

package main

import (
	"fmt"
	"strconv"
)

type List[T comparable] struct {
	next *List[T]
	prev *List[T]
	val T
}

func new[T comparable] (v T) *List[T]{
	return &List[T]{nil, nil, v}
}

func (l *List[T]) add(v T) {
	p := new(v)
	l.next.prev = p
	p.next = l.next 
	p.prev = l
	l.next = p
}

func (l *List[T]) delete (n T) bool {
	i := 0
	for p := l.next; p != l; p = p.next {
		if n == p.val {
			p.prev.next = p.next
			p.next.prev = p.prev 
			p.prev = nil 
			p.next = nil
			return true
		}
		i++
	}

	return false
}

func (l *List[T]) show() {
	fmt.Println("forward")
	for p := l.next; p != l; p = p.next {
		fmt.Printf("%v, ", p.val)
	}
	fmt.Println()

	fmt.Println("backward")
	for p := l.prev; p != l; p = p.prev {
		fmt.Printf("%v, ", p.val)
	}
	fmt.Println()
}

func main() {
	list := new("sentinel")			// sentinel
	list.prev = list
	list.next = list
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			list.add("hello" + strconv.Itoa(i))
		}
	}
	list.show()
	list.delete("hello1")
	list.show()
	list.delete("hello 3")
	list.show()
	list.delete("hello5")
	list.show()
}
