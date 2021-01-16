package main

import "fmt"

func main() {
	// BST Representation
	//         100
	//        /   \
	//      50    150
	//     / \    /  \
	//   20  80  110 200
	//  /  \
	// 10  30

	root := TreeNode{100, nil, nil}
	root.Left = &TreeNode{50, nil, nil}
	root.Right = &TreeNode{150, nil, nil}
	root.Left.Left = &TreeNode{20, nil, nil}
	root.Left.Left.Left = &TreeNode{10, nil, nil}
	root.Left.Left.Right = &TreeNode{30, nil, nil}
	root.Left.Right = &TreeNode{80, nil, nil}
	root.Right.Left = &TreeNode{110, nil, nil}
	root.Right.Right = &TreeNode{200, nil, nil}

	fmt.Println("BFS level order binary tree traversal", levelOrder(&root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	order := [][]int{}

	if root == nil {
		return order
	}

	nodeBFSQueue := []*TreeNode{}
	nodeBFSQueue = append(nodeBFSQueue, root)

	currentLevelNodes, nextLevelNodes := 1, 0

	levelOrderNodes := []int{}

	for len(nodeBFSQueue) > 0 {

		if currentLevelNodes > 0 {

			levelOrderNodes = append(levelOrderNodes, nodeBFSQueue[0].Val)

			if nodeBFSQueue[0].Left != nil {
				nodeBFSQueue = append(nodeBFSQueue, nodeBFSQueue[0].Left)
				nextLevelNodes++
			}

			if nodeBFSQueue[0].Right != nil {
				nodeBFSQueue = append(nodeBFSQueue, nodeBFSQueue[0].Right)
				nextLevelNodes++
			}

			nodeBFSQueue = nodeBFSQueue[1:]
			currentLevelNodes--
		}

		if currentLevelNodes == 0 {

			currentLevelNodes = nextLevelNodes
			nextLevelNodes = 0
			order = append(order, levelOrderNodes)
			levelOrderNodes = []int{}

		}

	}

	return order
}
