package main

func main() {

}

type MyQueue struct {
	items []int
}

func Constructor() MyQueue {
	test := MyQueue{}
	return test
}

func (this *MyQueue) Push(x int) {
	this.items = append(this.items, x)
}

func (this *MyQueue) Pop() int {
	removed := this.items[0]
	this.items = this.items[1:]
	return removed
}

func (this *MyQueue) Peek() int {
	if len(this.items) == 0 {
		return 0
	}
	return this.items[0]
}

func (this *MyQueue) Empty() bool {
	if len(this.items) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
