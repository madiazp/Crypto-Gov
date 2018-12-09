package block

import (
        "fmt"
        "crypto/sha256"
        tx "./TX"
)
///// definicion del bloque
type Block struct {

        prevHash        string
        //merkelTree      MTree
        timeStamp       string
        payload         tx.TX
        selfHash        string

}
func (u *Block) InitBlock(phash string, time string, load tx.TX){

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

func (u *Block) GetPrevHash() string{
        return u.prevHash
}

func (u *Block) GetHash() string{
        return u.selfHash
}
// funciones de utilidad

func (u *Block) BlockSize() int{
        return u.payload.TXSize()
}

func (u *Block) SignTheBlock(){
        content := toString(u.prevHash)+toString(u.timeStamp)+toString(u.payload) //+toString(u.merkelTree)
        u.selfHash = hexme(hashme(content))
}

////////////////////////////////////////////


///////////////////////////////////////////////////////////////
///////////////////Funciones auxiliares ///////////////////////
///////////////////////////////////////////////////////////////
func hexme( value interface{} ) string {
        str := toString(value)
        return hex.EncodeToString([]byte(str))

}

func hashme( value interface{} ) [32]byte{

        return sha256.Sum256([]byte(fmt.Sprintf("%v",value)))
}

func toString( value interface{} ) string{

        return fmt.Sprintf("%v",value)
}
