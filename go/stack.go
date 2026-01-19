package main

type StackNode struct {
	key  string
	prev *StackNode
}

type Stack struct {
	top *StackNode
}

func (this *Stack) push(key string) {
	if this.top == nil {
		this.top = &StackNode{key, nil}
	} else {
		node := StackNode{key, this.top}
		this.top = &node
	}
}

func (this *Stack) pop() string {
	if this.top == nil {
		return ""
	}
	ret := this.top.key
	this.top = this.top.prev
	return ret
}

func (this *Stack) print() string {
	a := ""
	node := this.top
	if node == nil {
		return a
	}
	for node.prev != nil {
		a += node.key + " "
		node = node.prev
	}
	a += node.key
	return a
}
