package p2p

import (
        "fmt"
        "crypto/sha256"
        "encoding/hex"
)

type Peer struct{
        addr        string
        port        string
        id          string

}
func (p *Peer) NewPeer(adr,prt string){
        p.addr = adr
        p.port = prt
        p.MakeId()
}
///// Metodos GET

func (p *Peer) GetAddr() string{
        return p.addr
}

func (p *Peer) GetPort() string{
        return p.port
}

func (p *Peer) GetID() string{
        return p.id
}

////// Utilidades
func (p *Peer) MakeId(){
        p.id = hexme(hashme(p.addr))
}


//////////////////////////////////////////////////////////
/////////////////////////Funciones Auxiliares //////////////
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
