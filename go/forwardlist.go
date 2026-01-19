package main

type ForwardListNode struct {
	key  string
	next *ForwardListNode
}

type ForwardList struct {
	head *ForwardListNode
}

func (this *ForwardList) insert(key string, index int) {
	if index == 0 {
		this.addHead(key)
		return
	}

	node := this.head
	for i := 0; i < index-1; i++ {
		if node == nil {
			break
		}
		node = node.next
	}
	if node == nil {
		return
	}
	newNode := &ForwardListNode{key, nil}
	newNode.next = node.next
	node.next = newNode
}

func (this *ForwardList) remove(index int) {
	if index == 0 {
		this.removeHead()
		return
	}

	node := this.head
	for i := 0; i < index-1; i++ {
		if node == nil {
			break
		}
		node = node.next
	}
	if node == nil {
		return
	}
	node.next = node.next.next
}

func (this *ForwardList) addTail(key string) {
	if this.head == nil {
		this.head = &ForwardListNode{key, nil}
		return
	}
	node := this.head
	for node.next != nil {
		node = node.next
	}
	node.next = &ForwardListNode{key, nil}
}

func (this *ForwardList) addHead(key string) {
	if this.head == nil {
		this.head = &ForwardListNode{key, nil}
		return
	}
	node := &ForwardListNode{key, nil}
	node.next = this.head
	this.head = node
}

func (this *ForwardList) removeTail() {
	if this.head == nil {
		return
	}

	node := this.head
	for node.next.next != nil {
		node = node.next
	}
	node.next = nil
}

func (this *ForwardList) removeHead() {
	if this.head == nil {
		return
	}

	this.head = this.head.next
}

func (this *ForwardList) printFromHead() string {
	node := this.head
	if node == nil {
		return ""
	}
	a := ""
	for node.next != nil {
		a += node.key + " "
		node = node.next
	}
	a += node.key
	return a
}

func printFromTailRec(s *string, node *ForwardListNode) {
	if node == nil {
		return
	}
	printFromTailRec(s, node.next)
	*s += node.key + " "
}

func (this *ForwardList) printFromTail() string {
	s := ""
	printFromTailRec(&s, this.head)
	return s
}

func (this *ForwardList) removeKey(key string, num int) {
	if num < 0 || this.head == nil {
		return
	}
	n := 0
	node := this.head
	if node.key == key && num == 0 {
		this.removeHead()
		return
	}
	for node.next != nil {
		if node.next.key == key {
			if n == num {
				node.next = node.next.next
				return
			}
			n++
		}
		node = node.next
	}
}

func (this *ForwardList) find(key string, num int) *ForwardListNode {
	if num < 0 {
		return nil
	}
	n := 0
	node := this.head
	for node != nil {
		if node.key == key {
			if n == num {
				return node
			}
			n++
		}
		node = node.next
	}
	return nil
}
