package AuxStruct

import (
        "fmt"
        "crypto/sha256"
)
/////////////////////////////////////////
//Merkle Tree

type MTree struct {
        root *Node
        hashroot []byte
        midnodes []*Node
        leafs []*Node
}

func (m *MTree) InitTree(load TX) {

}
func (m *MTree) GetHashroot() []byte {

}
func (m *MTree) GetRoot() *Node {

}
func (m *MTree) GetLeafs() []*Node{

}
func (m *MTree) GetALeaf( leafhash []byte ) *Node{

}
func (m *MTree) GetANode (nodehash []byte) *Node{

}

//////////////////////////////////////
// Node
type Node struct {
        parent *Node
        lchild *Node
        rchild *Node
        leaf bool
        dup bool
        hash []byte
}

func (n *Node) InitNode() {

}

func (n *Node) GetHash() []byte {

}

func (n *Node) IsALeaf() bool {

}

func (n *Node) IsDuplicated() bool {

}

func (n *Node) SetLeafHash( leafhash []byte){

}

func (n *Node) SelfVerify() bool{

}
