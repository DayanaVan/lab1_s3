package main

type QueueNode struct {
	key  string
	next *QueueNode
}

type Queue struct {
	top *QueueNode
	bot *QueueNode
}

func (this *Queue) push(key string) {
	newNode := &QueueNode{key, nil}
	if this.top == nil {
		this.bot = newNode
		this.top = newNode
		return
	} else {
		this.top.next = newNode
		this.top = newNode
	}
}

func (this *Queue) pop() string {
	if this.bot == nil {
		return ""
	}

	node := this.bot
	this.bot = node.next
	if this.bot == nil {
		this.top = nil
	}
	return node.key
}

func (this *Queue) print() string {
	node := this.bot
	if node == nil {
		return ""
	}

	a := ""
	for node != nil {
		a += node.key + " "
		node = node.next
	}
	return a
}
