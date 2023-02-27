package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 1.哈希链表
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	p := head
	ans := head
	for head != nil {
		m[head] = &Node{
			Val:    head.Val,
			Next:   nil,
			Random: nil,
		}
		head = head.Next
	}
	for p != nil {
		if p.Next != nil {
			m[p].Next = m[p.Next]
		}
		if p.Random != nil {
			m[p].Random = m[p.Random]
		}
		p = p.Next
	}
	return m[an]
}

// 2.官方哈希链表
var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{
		Val:    node.Val,
		Next:   nil,
		Random: nil,
	}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}

// 3.常数空间 迭代法 难想到
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
		nextNode := node.Next.Next
		if node.Next.Next != nil {
			node.Next.Next = node.Next.Next.Next
		}
		node = nextNode
	}
	headNew := head.Next
	return headNew
}

func main() {
	copyRandomList(&Node{Val: 7, Next: &Node{Val: 13, Next: &Node{Val: 11, Next: &Node{Val: 10, Next: &Node{Val: 1}}}}, Random: &Node{Val: 11, Next: &Node{Val: 10, Next: &Node{Val: 1}}}})
}
