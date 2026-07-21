
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)
type Node [T any] struct{
	val T
	next *Node[T]
}
type LinkedList[T any]struct{
	head *Node[T]
}
func (l *LinkedList[T]) Prepend(i T) {
	cur := &Node[T]{}
	cur.val = i
	cur.next = l.head
	l.head = cur
}

func (l *LinkedList[T]) Append(i T) {
	z := &Node[T]{}
	z.val = i
	z.next = nil
	if l.head == nil{
		l.head = z
		return
	}
	cur := l.head
	for cur.next != nil{
		cur = cur.next
	}
	cur.next = z

}

func (l *Node[T]) MarshalJSON() ([]byte, error) {
	var NextN any //any чтобы было пусто
	if l ==nil{
		return json.Marshal(nil)
	}
	if l.next != nil{
		NextN = l.next //
	}
 	return json.Marshal(struct {
  	Value     T    `json:"value,omitempty"`
  	Next   any `json:"child,omitempty"`
 	}{
  	Value:     l.val,   
  	Next:   NextN, 

 })
}
func (n *Node[T]) UnmarshalJSON(data []byte ) (error){
 type NodeDTO struct{
	Value json.RawMessage `json:"value"`
	Child *json.RawMessage `json:"child,omitempty"`
	}
	var i NodeDTO
 	err := json.Unmarshal(data, &i)
 	if err != nil{
		return err
 	}
 	err2 := json.Unmarshal(i.Value, &n.val)
 	if err2 != nil{
		return err2
 	}
 	if i.Child!= nil{
		var j Node[T]
		n.next = &j
		err3 := json.Unmarshal(*i.Child, n.next)
		if err3 != nil{
			return err3
		}
	}
 	return nil
}	

func (l *LinkedList[T]) UnmarshalJSON(data []byte) (error){
	type ListDTO struct{
		Head *json.RawMessage `json:"head,omitempty"`
		Lenght int `json:"lenght,omitempty"`
	}
	var i ListDTO
	err2 := json.Unmarshal(data, &i)
	if err2 != nil{
		return err2
	}
	if i.Head != nil{
		var j Node[T]
		l.head = &j
		err3 := json.Unmarshal(*i.Head, l.head)
		if err3 != nil{
			return err3
		}
	}
	return nil
}


func (t *LinkedList[T]) MarshalJSON() ([]byte, error) {
	if t ==nil{
		return json.Marshal(nil)
	}
	Leng :=t.Len()
 	return json.Marshal(struct {
		Head   *Node[T] `json:"head,omitempty"`
		Length int `json:"length"`
 	}{
  		Head:     t.head,
		Length: Leng,


 })
}
func (l *LinkedList[T]) Len() int{
	c  := 0
	cur := l.head
	for cur != nil{
		c++
		cur = cur.next
	}
	return c
}
func (l *LinkedList[T]) Print(){
	cur := l.head
	for cur != nil{
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
	for i := range words{
		list.Append(words[i])
	}

	data, err := json.MarshalIndent(list, "", " ")
	if err != nil{
		fmt.Println(err)
	}
	var list2 LinkedList[string]
	err2 := json.Unmarshal(data, &list2)
	if err2 !=nil{
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
