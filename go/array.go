package main

type Array struct {
	memorySize int
	size       int
	memory     []string
}

func newArray() *Array {
	arr := Array{0, 0, nil}
	arr.memory = make([]string, 0)
	return &arr
}

func (this *Array) add(key string) {
	if this.memorySize < this.size+1 {
		oldMemory := this.memory
		if this.size == 0 {
			this.memorySize = 1
		} else {
			this.memorySize = this.size * 2
		}
		this.memory = make([]string, this.memorySize)
		for i := 0; i < this.size; i++ {
			this.memory[i] = oldMemory[i]
		}
	}
	this.memory[this.size] = key
	this.size++
}

func (this *Array) insert(key string, index int) {
	if index > this.size || index < 0 {
		return
	}

	if this.memorySize < this.size+1 {
		oldMemory := this.memory
		if this.size == 0 {
			this.memorySize = 1
		} else {
			this.memorySize = this.size * 2
		}
		this.memory = make([]string, this.memorySize)
		for i := 0; i < index; i++ {
			this.memory[i] = oldMemory[i]
		}
		this.memory[index] = key
		for i := index + 1; i < this.size+1; i++ {
			this.memory[i] = oldMemory[i-1]
		}
	} else {
		for i := this.size; i > index; i-- {
			this.memory[i] = this.memory[i-1]
		}
		this.memory[index] = key
	}
	this.size++
}

func (this *Array) get(index int) string {
	if index >= this.size || index < 0 {
		return ""
	}
	return this.memory[index]
}

func (this *Array) remove(index int) {
	if index >= this.size || index < 0 {
		return
	}

	for i := index; i < this.size-1; i++ {
		this.memory[i] = this.memory[i+1]
	}
	this.size--
}

func (this *Array) change(index int, key string) {
	if index >= this.size || index < 0 {
		return
	}
	this.memory[index] = key
}

func (this *Array) print() string {
	if this.size == 0 {
		return ""
	}

	a := ""
	for i := 0; i < this.size-1; i++ {
		a += this.memory[i] + " "
	}
	a += this.memory[this.size-1]
	return a
}
