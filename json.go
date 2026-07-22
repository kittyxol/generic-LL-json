package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Node[T any] struct {
	val  T
	next *Node[T]
}

func (n *Node[T]) MarshalJSON() ([]byte, error) {
	if n == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(struct {
		Value T        `json:"value,omitempty"`
		Next  *Node[T] `json:"child,omitempty"`
	}{
		Value: n.val,
		Next:  n.next,
	})
}

func (n *Node[T]) UnmarshalJSON(data []byte) error {
	type nodeDTO struct { //type nodeDTO structure. Для JSON важны публичные поля, структура может быть приватной
		Value T        `json:"value"`
		Child *nodeDTO `json:"child"`
	}
	var i nodeDTO
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	n.val = i.Value
	curN := n                 //корень
	curDTO := &i              //дто
	for curDTO.Child != nil { //
		curN.next = &Node[T]{}
		curN = curN.next
		curDTO = curDTO.Child
		curN.val = curDTO.Value
	}
	return nil
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Prepend(i T) {
	l.head = &Node[T]{val: i, next: l.head}
}

func (l *LinkedList[T]) Append(i T) {
	z := &Node[T]{val: i, next: nil}
	if l.head == nil {
		l.head = z
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = z

}

func (l *LinkedList[T]) UnmarshalJSON(data []byte) error {
	type listDTO struct {
		Head   *Node[T] `json:"head"`
		Lenght int      `json:"lenght"`
	}
	var i listDTO
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	if i.Head != nil {
		l.head = i.Head
	}
	return nil
}

func (l *LinkedList[T]) MarshalJSON() ([]byte, error) {
	if l == nil {
		return json.Marshal(nil)
	}
	Leng := l.Len()
	return json.Marshal(struct {
		Head   *Node[T] `json:"head,omitempty"`
		Length int      `json:"length"`
	}{
		Head:   l.head,
		Length: Leng,
	})
}
func (l *LinkedList[T]) Len() int {
	c := 0
	cur := l.head
	for cur != nil {
		c++
		cur = cur.next
	}
	return c
}
func (l *LinkedList[T]) Print() {
	cur := l.head
	for cur != nil {
		fmt.Printf("%v ->", cur.val)
		cur = cur.next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList[string]{}
	fmt.Println("Введите строку:")
	scanner := bufio.NewReader(os.Stdin)
	text, _ := scanner.ReadString('\n')
	words := strings.Fields(text)
	for i := range words {
		list.Append(words[i])
	}

	data, err := json.MarshalIndent(list, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	var list2 LinkedList[string]
	err2 := json.Unmarshal(data, &list2)
	if err2 != nil {
		fmt.Println(err2)
	}
	list2.Print()
	fmt.Println(string(data))
	// list := LinkedList{...}
	// b, _ := json.Marshal(list)
	// var newList LinkedList
	// json.Unmarshal(&b, &newList)
	// fmt.Println(newList)

}
