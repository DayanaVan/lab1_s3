all: l1

l1: main.o array.o forwardlist.o io.o list.o queue.o bstree.o stack.o
	g++ main.o array.o forwardlist.o io.o list.o queue.o bstree.o stack.o -o l1

main.o: main.cpp
	g++ main.cpp -c -o main.o

array.o: array.cpp array.h
	g++ array.cpp -c -o array.o

forwardlist.o: forwardlist.cpp forwardlist.h
	g++ forwardlist.cpp -c -o forwardlist.o

list.o: list.cpp list.h
	g++ list.cpp -c -o list.o

queue.o: queue.cpp queue.h
	g++ queue.cpp -c -o queue.o

bstree.o: bstree.cpp bstree.h
	g++ bstree.cpp -c -o bstree.o

io.o: io.cpp io.h
	g++ io.cpp -c -o io.o

stack.o: stack.cpp stack.h
	g++ stack.cpp -c -o stack.o
