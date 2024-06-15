//Let say you have a tree. you need to return a list which will contain all the  root to leaf path sum.
//                1
//        2               3

// 4            5    6         7

//result:= [7,8,10,11]

package main

import (
    "fmt"
)


type Node struct {
    data int
    left *Node
    right *Node
}

func newNode(input int) *Node {
    return &Node{
        data: input,
    }
}


func RootToLeafSum(node *Node, sum int) []int {
    result := make([]int, 0)
    if node == nil {
        return result
    } else {
        sum += node.data
        
        if node.left == nil && node.right == nil {
            result = append(result, sum)
        }
        
        
        
        leftsum := RootToLeafSum(node.left, sum)
        result = append(result, leftsum...)
        
        
        rightSum := RootToLeafSum(node.right, sum)
        result = append(result, rightSum...)
        
    }
    
    return result
}

func RootToLeafSumWithoutReturn(node *Node, sum int, result *[]int) {
    if node == nil {
        return
    } else {
        sum += node.data
        
        if node.left == nil && node.right == nil {
            *result = append(*result, sum)  //remember this to modify the passed slice 
        }
        
        
        RootToLeafSum(node.left, sum, result)
        
        
        RootToLeafSum(node.right, sum, result)
        
    }
    return 
}

func main() {
    
    root := newNode(1)
    
    root.left = newNode(2)
    root.right = newNode(3)
    
    root.left.left = newNode(4)
    root.left.right = newNode(5)
    
    root.right.left = newNode(6)
    root.right.right = newNode(7)
    
    
    result := RootToLeafSum(root, 0)
    
    result2 := make([]int, 0)
    RootToLeafSumWithoutReturn(root, 0, result2)
    
    fmt.Println(result)
    
    fmt.Println(result2)
}


