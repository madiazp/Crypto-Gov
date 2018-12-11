package p2p

import (
        "fmt"
        "crypto/sha256"
        "encoding/hex"
)

type Peer struct{
        Addr        string
        Port        string
        Id          string

}
func (p *Peer) NewPeer(adr,prt string){
        fmt.Println(adr)
        fmt.Println(prt)
        p.Addr = adr
        p.Port = prt
        p.MakeId()
}
///// Metodos GET

func (p *Peer) GetAddr() string{
        return p.Addr
}

func (p *Peer) GetPort() string{
        return p.Port
}

func (p *Peer) GetID() string{
        return p.Id
}

////// Utilidades
func (p *Peer) MakeId(){
        p.Id = hexme(hashme(p.Addr))
}


//////////////////////////////////////////////////////
//////////////////////////////////////////////////////

type PPeer struct{
        *Peer
        parent          Peer
        children        map[string]Peer
        leafs           map[string]Peer
        rtcache         map[string]Peer
}
////////// Metodos Get


func (p *PPeer) GetLeafs() map[string]Peer{
        return p.leafs
}

func (p *PPeer) GetRoutingTableCache() map[string]Peer{
        return p.rtcache
}

/////////// Metodos de configuracion de la red

func (p *PPeer) NewPPeer(adr, prt string){

        p.NewPeer(adr,prt)
        //p.parent = prnt


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
