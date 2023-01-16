package main

import (
	"fmt"
)

type Container interface{
  begin() Iterator
  append(v interface{})
}

type Iterator interface{
  Next() Iterator
  IsEnd() bool
  Deref() interface{}
}

type Node struct {
	next *Node
	value  interface{}
}

type List struct {
	head *Node
	i int
}

type listItrable struct{
  curr *Node
}

type Vector struct{
  vectorInt []interface{}
}

type vectorItrable struct{
  i int
  vectorInt []interface{}
}

func (l *List) begin() Iterator{
  return &listItrable{l.head}
}

func (l *listItrable) IsEnd() bool{
  if (l.curr != nil){
    return false
  }
  return true
}

func (l *listItrable) Next() Iterator {
  return &listItrable{l.curr.next}
}

func (l *listItrable) Deref() interface{} {
  return l.curr.value
}

func (l *List) append(v interface{}) {
  if l.head == nil {
    l.head = new(Node)
    l.head.value = v
    l.head.next = nil
  } else {
    nike := l.head
    for nike.next != nil {
      nike = nike.next
    }
    nike.next = new(Node)
    nike.next.value = v
    nike.next.next = nil
  }
}

func (u *Vector) begin() Iterator {
  return &vectorItrable{
    vectorInt: u.vectorInt,
  }
}

func (u *vectorItrable) IsEnd() bool {
  if u.i >= len(u.vectorInt) {
    return true
  }
  return false
}

func(u *vectorItrable) Deref() interface{} {
  return u.vectorInt[u.i]
}

func(u *vectorItrable) Next() Iterator {
  return &vectorItrable{
    u.i+1,u.vectorInt}
}

func(u *Vector) append(v interface{}) {
  u.vectorInt = append(u.vectorInt, v)
}

func SumInt (c Container) int {
  var total int
  i := c.begin()
  for i.IsEnd() != true {
    switch i.Deref().(type) {
      case int:
      total += i.Deref().(int)
    }
    i = i.Next()
  }
  return total
}

func SumFloat64 (c Container) float64 {
  var total float64
  i := c.begin()
  for i.IsEnd() != true {
    switch i.Deref().(type) {
      case float64:
      total += i.Deref().(float64)
    }
    i = i.Next()
  }
  return total
}

func main() {
	fmt.Println("Hello, World!")
  var v Vector
  v.append(10)
  v.append(20)
  v.append(30)
  var v2 Iterator
  v2 = v.begin()
  fmt.Println(v2.Deref())
  v2 = v2.Next()
  fmt.Println(v2.Deref())
  v2 = v2.Next()
  fmt.Println(v2.Deref())
  fmt.Println(v.begin().Next().Deref())
  var L List
  var fs Container = &L
  L.append(0.75)
  L.append(100)
  L.append(2)
  L.append(3)
  L.append(0.25)
  fmt.Printf("sum of integer: %d\n", SumInt(fs))
  fmt.Printf("sum of integer: %f\n", SumFloat64(fs))
}