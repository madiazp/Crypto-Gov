package block

import (
        "fmt"
        "crypto/sha256"
)
///// definicion del bloque
type Ublock struct {

      prevHash []byte
      merkelTree MTree
      timeStamp string
      payload TX

}
func (u *Ublock) InitBlock(phash []byte, time string, load TX){
        u.prevHash = phash
        u.timeStamp = time
        u.payload = load
        u.merkelTree.InitTree(load)

}

func (u *Ublock) GetMSGS() []string{
        return payload.GetMsgs()
}

func (u *Ublock) VerifyTXS() (bool){
        return u.payload.VerifyTXS()
}

func (u *Ublock) VerifySignTXS() bool{

}

/////////////////////////////////////

type Block struct {
      ublk Ublock
      selfhash []byte
}
func (b *Block) InitBlock( blk Ublock){
        b.ublk = blk
        selfhash = hashme(blk)
}
func (b *Block) GetMsgs() []String {
      return b.ublk.GetMSGS()
}

func (b *Block) VerifyTXS() (bool, error)  {
    if !b.VerifyTXS() {
            return false
    }

    return b.SelfVerifySign()

}

func (b *Block) SelfSign() error {
      var err error
      b.selfhash, err = sha256.Sum256(fmt.Sprintf("%v",b.ublk))
      return err
}

func (b *Block) SelfVerifySign() (bool, error) {
      test, err := sha256.Sum256(fmt.Sprintf("%v",b.ublk))
      if err != nil || test != b.selfhash{
        return false, err
      }
      return true
}
func (b *Block) SelfVerify() (bool, error) {
      if (b.SelfVerifySign()) { return false,nil }

}

////////////////////////////////////////////7

////////////////////////////////////////////////////////////////////////////////
///////////////////Estructuras Auxiliares //////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
//Transacciones
type TX struct {

      msgs []string
}

func (t *TX) GetMsgs() []string {
      return t.msgs
}

func (t *TX) AddMsg (mg string) {
      append(t.msgs,mg)
}

func (t *TX) VerifyMSGS() (bool)  {
      for i := range t.msgs {
        if t.msgs[i] != "True" {
          return false
        }
      }
      return true
}


/////////////////////////////////////////
//Merkle Tree

type MTree struct {
        root *Node
        hashroot []byte
        midnodes []*Node
        leafs []*Node
}

func (m *MTree) InitTree(load TX) {
        msgs := load.GetMsgs()
        for _,msg := range msgs {

        }
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
///////////////////////////////////////////////////////////////
///////////////////Funciones auxiliares ///////////////////////
///////////////////////////////////////////////////////////////


func hashme( value interface{} ) []byte{
        return sha256.Sum256([]byte(fmt.Sprintf("%v",value)))
}
