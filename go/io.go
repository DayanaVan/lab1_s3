package main

import (
	"bufio"
)

func readArray(scanner *bufio.Scanner, arr *Array) {
	for scanner.Scan() {
		arr.add(scanner.Text())
	}
}

func writeArray(writer *bufio.Writer, arr *Array) {
	writer.WriteString(arr.print())
	writer.Flush()
}

func readForwardList(scanner *bufio.Scanner, fl *ForwardList) {
	var tail *ForwardListNode
	for scanner.Scan() {
		if fl.head == nil {
			fl.addHead(scanner.Text())
			tail = fl.head
		} else {
			tail.next = &ForwardListNode{scanner.Text(), nil}
			tail = tail.next
		}
	}
}

func writeForwardList(writer *bufio.Writer, fl *ForwardList) {
	writer.WriteString(fl.printFromHead())
	writer.Flush()
}

func readList(scanner *bufio.Scanner, l *List) {
	for scanner.Scan() {
		l.addTail(scanner.Text())
	}
}

func writeList(writer *bufio.Writer, l *List) {
	writer.WriteString(l.printFromHead())
	writer.Flush()
}

func readStack(scanner *bufio.Scanner, s *Stack) {
	for scanner.Scan() {
		s.push(scanner.Text())
	}
}

func writeStackRec(writer *bufio.Writer, sn *StackNode) {
	if sn == nil {
		return
	}
	writeStackRec(writer, sn.prev)
	writer.WriteString(sn.key + " ")
}

func writeStack(writer *bufio.Writer, s *Stack) {
	writeStackRec(writer, s.top)
	writer.Flush()
}

func readQueue(scanner *bufio.Scanner, q *Queue) int {
	for scanner.Scan() {
		q.push(scanner.Text())
	}
	return 0
}

func writeQueue(writer *bufio.Writer, q *Queue) {
	writer.WriteString(q.print())
	writer.Flush()
}
