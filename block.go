package block

import (
        "fmt"
        "crypto/sha256"
        tx "./TX"
)
///// definicion del bloque
type Block struct {

        prevHash        [32]byte
        //merkelTree      MTree
        timeStamp       string
        payload         tx.TX
        selfHash        [32]byte

}
func (u *Block) InitBlock(phash [32]byte, time string, load tx.TX){

        u.prevHash = phash
        u.timeStamp = time
        u.payload = load
        //u.merkelTree.InitTree(load)
        u.SignTheBlock()

}
// funciones GET

func (u *Block) GetContent() tx.TX{
        return u.payload
}

func (u *Block) GetTimeStamp() string{
        return u.timeStamp
}

func (u *Block) GetPrevHash() [32]byte{
        return u.prevHash
}

func (u *Block) GetHash() [32]byte{
        return u.selfHash
}
// funciones de utilidad
func (u *Block) BlockSize() int{
        return u.payload.TXSize()
}

func (u *Block) SignTheBlock(){
        content := toString(u.prevHash)+toString(u.timeStamp)+toString(u.payload) //+toString(u.merkelTree)
        u.selfHash = hashme(content)
}

////////////////////////////////////////////


///////////////////////////////////////////////////////////////
///////////////////Funciones auxiliares ///////////////////////
///////////////////////////////////////////////////////////////


func hashme( value interface{} ) [32]byte{

        return sha256.Sum256([]byte(fmt.Sprintf("%v",value)))
}

func toString( value interface{} ) string{

        return fmt.Sprintf("%v",value)
}
