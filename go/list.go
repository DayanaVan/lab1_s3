package main

type ListNode struct {
	key  string
	next *ListNode
	prev *ListNode
}

type List struct {
	tail *ListNode
	head *ListNode
}

func (this *List) insert(key string, index int) {
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
	if node.next == this.tail {
		this.addTail(key)
		return
	}
	newNode := &ListNode{key, nil, nil}
	newNode.next = node.next
	newNode.prev = node
	node.next = newNode
	newNode.next.prev = node.prev
}

func (this *List) remove(index int) {
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
	if node.next == this.tail {
		this.removeTail()
		return
	}
	toDelete := node.next
	if toDelete != nil {
		node.next = toDelete.next
		toDelete.next.prev = node
	}
}

func (this *List) addTail(key string) {
	node := &ListNode{key, nil, nil}
	if this.tail == nil {
		this.tail = node
		this.head = node
		return
	}
	node.prev = this.tail
	this.tail.next = node
	this.tail = node
}

func (this *List) addHead(key string) {
	node := &ListNode{key, nil, nil}
	if this.tail == nil {
		this.tail = node
		this.head = node
		return
	}
	node.next = this.head
	this.head.prev = node
	this.head = node
}

func (this *List) removeTail() {
	if this.tail == nil {
		return
	}
	if this.tail == this.head {
		this.tail = nil
		this.head = nil
		return
	}
	this.tail = this.tail.prev
	this.tail.next = nil
}

func (this *List) removeHead() {
	if this.head == nil {
		return
	}
	if this.tail == this.head {
		this.tail = nil
		this.head = nil
		return
	}
	this.head = this.head.next
	this.head.prev = nil
}

func (this *List) printFromHead() string {
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

func (this *List) printFromTail() string {
	node := this.tail
	if node == nil {
		return ""
	}
	a := ""
	for node.prev != nil {
		a += node.key + " "
		node = node.prev
	}
	a += node.key
	return a
}

func (this *List) removeKey(key string, num int) {
	if num < 0 {
		return
	}
	n := 0
	node := this.head
	for node != nil {
		if node.key == key {
			if n == num {
				if node == this.tail {
					this.removeTail()
				} else if node == this.head {
					this.removeHead()
				} else {
					node.prev.next = node.next
					node.next.prev = node.prev
					node.next = node.next.next
				}
				return
			}
			n++
		}
		node = node.next
	}
}

func (this *List) find(key string, num int) *ListNode {
	if num < 0 {
		return nil
	}
	n := 0
	node := this.tail
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
