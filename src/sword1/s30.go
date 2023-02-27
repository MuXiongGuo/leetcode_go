package main

type MinStack struct {
	minStk, data []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minStk) == 0 {
		this.minStk = append(this.minStk, x)
	} else {
		this.minStk = append(this.minStk, min(this.minStk[len(this.minStk)-1], x))
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minStk = this.minStk[:len(this.minStk)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	return this.minStk[len(this.minStk)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
