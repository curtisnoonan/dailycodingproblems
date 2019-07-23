package main

import (
	"fmt"
	"errors"
	"log"
	"strconv"
)

type Node struct {
	self int
	val  int
	left  *Node
	right *Node
}
type Tree struct {
	Root *Node
}
func (t *Tree) Insert(value, data int) error {
	if t.Root == nil {
		t.Root = &Node{self: value, val: data}
		return nil
	}
	return t.Root.Insert(value, data)
}
func (n *Node) Insert(value, data int) error{
	if n == nil {
		return errors.New("Cannot insert value into nil tree")
	}
	switch{
	case value == n.val:
		return nil
	case value < n.val:
		if n.left == nil {
			n.left = &Node{self:value, val:data}
			return nil
		}
		return n.left.Insert(value,data)
	case value > n.val:
		if n.right == nil{
			n.right = &Node{self:value,val:data}
			return nil
		}
		return n.right.Insert(value,data)
	}
	return nil	
}
func (t *Tree) Serialize(n *Node) string {
	serialized := ""
	if n == nil {
		return ""
	}
	t.Serialize(n.left)
	t.Serialize(n.right)
	serialized += strconv.Itoa(n.val)
	return serialized
}

func main() {
	// Set up a slice of strings.
	values := []int{0,1,2,3,4,5}
	data := []int{20,30,10,15,20,34}

	// Create a tree and fill it from the values.
	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}
	serialized := tree.Serialize(tree.Root)
	fmt.Println(serialized)
}