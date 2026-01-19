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

func (this *ForwardList) addBefore(key, targetKey string, num int) {
	if this.head == nil {
		return
	}
	
	if this.head.key == targetKey {
		count := 0
		node := this.head
		for node != nil && node.key == targetKey {
			if count == num {
				newNode := &ForwardListNode{key, nil}
				newNode.next = this.head
				this.head = newNode
				return
			}
			count++
			node = node.next
		}
	}

	count := 0
	prev := this.head
	for prev.next != nil {
		if prev.next.key == targetKey {
			if count == num {
				newNode := &ForwardListNode{key, nil}
				newNode.next = prev.next
				prev.next = newNode
				return
			}
			count++
		}
		prev = prev.next
	}
}

func (this *ForwardList) addAfter(key, targetKey string, num int) {
	if this.head == nil {
		return
	}
	
	count := 0
	node := this.head
	for node != nil {
		if node.key == targetKey {
			if count == num {
				newNode := &ForwardListNode{key, nil}
				newNode.next = node.next
				node.next = newNode
				return
			}
			count++
		}
		node = node.next
	}
}

func (this *ForwardList) removeBefore(targetKey string, num int) {
	if this.head == nil || this.head.next == nil {
		return
	}

	count := 0
	if this.head.next.key == targetKey {
		temp := this.head
		for temp != nil && temp.key == this.head.next.key {
			if count == num {
				this.removeHead()
				return
			}
			count++
			temp = temp.next
		}
	}

	count = 0
	prev := this.head
	for prev.next != nil && prev.next.next != nil {
		if prev.next.next.key == targetKey {
			if count == num {
				prev.next = prev.next.next
				return
			}
			count++
		}
		prev = prev.next
	}
}

func (this *ForwardList) removeAfter(targetKey string, num int) {
	if this.head == nil {
		return
	}
	
	count := 0
	node := this.head
	for node != nil && node.next != nil {
		if node.key == targetKey {
			if count == num {
				node.next = node.next.next
				return
			}
			count++
		}
		node = node.next
	}
}
