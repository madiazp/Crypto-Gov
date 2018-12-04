package block

import (
        "fmt"
        "crypto/sha256"
)
///// definicion del bloque
type Block struct {

        prevHash []byte
        merkelTree MTree
        timeStamp string
        payload TX
        selfHash []byte

}
func (u *Block) InitBlock(phash []byte, time string, load TX){

        u.prevHash = phash
        u.timeStamp = time
        u.payload = load
        u.merkelTree.InitTree(load)
        u.SignTheBlock()

}

func (u *Block) GetContent() map[[]byte]{
        return payload.GetMsgs()
}

func (u *Block) GetAContent( h []byte) string{

}

func (u *Block) VerifyTXS() (bool){
        return u.payload.VerifyTXS()
}

func (u *Block) VerifySignTXS() bool{

}
func (u *Block) SignTheBlock(){
        content := toString(u.prevHash)+toString(u.timeStamp)+toString(u.payload)+toString(u.merkelTree)
        u.selfHash = hashme(content)
}
////////////////////////////////////////////


///////////////////////////////////////////////////////////////
///////////////////Funciones auxiliares ///////////////////////
///////////////////////////////////////////////////////////////


func hashme( value interface{} ) []byte{

        return sha256.Sum256([]byte(fmt.Sprintf("%v",value)))
}

func toString( value interface{} ) string{

        return fmt.Sprintf("%v",value)
}
