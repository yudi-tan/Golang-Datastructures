package main

import "fmt"

//BST binary search tree struct
type BST struct {
	value int
	left  *BST
	right *BST
}

func (bst *BST) insert(item int) {
	if item < bst.value {
		if bst.left == nil {
			bst.left = &BST{value: item, left: nil, right: nil}
		} else {
			bst.left.insert(item)
		}
	} else {
		if bst.right == nil {
			bst.right = &BST{value: item, left: nil, right: nil}
		} else {
			bst.right.insert(item)
		}
	}
}

func (bst *BST) contains(item int) bool {
	if bst.value == item {
		return true
	} else if bst.value != item && bst.left == nil && bst.right == nil {
		return false
	} else {
		return bst.left.contains(item) || bst.right.contains(item)
	}
}

func (bst *BST) getMin() int {
	if bst.left == nil {
		return bst.value
	} else {
		return bst.left.getMin()
	}
}

func (bst *BST) getMax() int {
	if bst.right == nil {
		return bst.value
	} else {
		return bst.right.getMax()
	}
}

func (bst *BST) printBST() {
	if bst.left == nil && bst.right == nil {
		fmt.Print(bst.value)
	} else {
		fmt.Print(bst.value)
		bst.left.printBST()
		bst.right.printBST()
	}
}

func main() {
	test := BST{value: 5, left: nil, right: nil}
	test.left = &BST{value: 3, left: nil, right: nil}
	test.left.left = &BST{value: 2, left: nil, right: nil}
	test.left.right = &BST{value: 4, left: nil, right: nil}
	test.right = &BST{value: 7, left: nil, right: nil}
	test.right.left = &BST{value: 6, left: nil, right: nil}
	test.right.right = &BST{value: 8, left: nil, right: nil}
	test.printBST()
	fmt.Println("")
	test2 := BST{value: 5, left: nil, right: nil}
	test2.insert(3)
	test2.insert(2)
	test2.insert(4)
	test2.insert(7)
	test2.insert(6)
	test2.insert(8)
	test.printBST()
	fmt.Println("")
	fmt.Print(test2.contains(4))
	fmt.Println("")
	fmt.Print(test2.getMin())
	fmt.Println("")
	fmt.Print(test2.getMax())
}
