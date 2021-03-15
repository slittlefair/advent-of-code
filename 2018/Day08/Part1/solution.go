package main

import (
	"Advent-of-Code"
	"fmt"
	"strings"
)

type node struct {
	id                 int
	childNodesQuantity int
	metaDataQuantity   int
	childNodes         []node
	metaData           []int
}

var tree []int
var metaDataTotal = 0

func createNodes(i int) node {
	n := node{
		id:                 i,
		childNodesQuantity: tree[i],
		metaDataQuantity:   tree[i+1],
	}
	i += 2
	for j := 0; j < n.childNodesQuantity; j++ {
		childNode := createNodes(i)
		n.childNodes = append(n.childNodes, childNode)
		i += nodeLength(childNode)
	}
	for j := 0; j < n.metaDataQuantity; j++ {
		n.metaData = append(n.metaData, tree[i+j])
		metaDataTotal += tree[i+j]
	}
	return n
}

func nodeLength(n node) int {
	l := 2
	for _, v := range n.childNodes {
		l += nodeLength(v)
	}
	l += n.metaDataQuantity
	return l
}

func main() {
	oldTree := helpers.ReadFile()
	oldTree = strings.Split(oldTree[0], " ")
	for _, v := range oldTree {
		tree = append(tree, helpers.StringToInt(v))
	}
	createNodes(0)
	fmt.Println(metaDataTotal)
}
