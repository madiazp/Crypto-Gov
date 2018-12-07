package peer

import (
        "fmt"
        "crypto/sha256"
)

type Peer struct{
        addr        string
        port        string
        id      [32]byte
        leafs       map[string]Peer

}
///// Metodos GET

func (p *Peer) GetAddr() string{
        return p.addr
}

func (p *Peer) GetPort() string{
        return p.port
}

func (p *Peer) GetID() [32]byte{
        return p.id
}

func (p *Peer) GetLeafs() map[string]Peer{
        return p.leafs
}

func (p *Peer) GetALeaf( pid string) Peer{
        return p.leafs(pid)
}

func (p *Peer) GetLeafsIds() []string{
        var lfsid []string
        for _,ids := range p.leafs{
                append(ids,lfsid)
        }
        return lfsid
}

func (p *Peer) GetAmountLeafs() int {
        return len(p.leafs)
}
//// Metodos de utilidad

func (p *Peer) AddALeaf( leaf Peer) {
        p.leafs[toString(leaf.id)] = leaf

}

func (p *Peer) MakeId(){
        p.id = hashme(p.addr)
}


//////////////////////////////////////////////////////////
/////////////////////////Funciones Auxiliares //////////////

func hashme( value interface{} ) [32]byte{

        return sha256.Sum256([]byte(fmt.Sprintf("%v",value)))
}

func toString( value interface{} ) string{

        return fmt.Sprintf("%v",value)
}
