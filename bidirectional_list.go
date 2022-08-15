// 双方向リスト, 挿入, 削除, 照会 のジェネリクス

package main

import (
	"fmt"
)

type List[T any] struct {
	next *List[T]
	prev *List[T]
	val T
}

func new[T any] (v T) *List[T]{
	return &List[T]{nil, nil, v}
}

func (l *List[T]) add(v T) {
	p := new(v)
	l.next.prev = p
	p.next = l.next 
	p.prev = l
	l.next = p
}

func (l *List[T]) delete (n int) bool {
	i := 0
	for p := l; p != l.prev; p = p.next {
		if n == i && p.prev != nil {
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
	list := new(0)			// sentinel
	list.prev = list
	list.next = list
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			list.add(i)
		}
	}
	list.show()
	list.delete(1)
	list.show()
	list.delete(2)
	list.show()
	list.delete(3)
	list.show()
}
