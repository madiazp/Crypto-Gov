package p2p

import (
        "fmt"
        "crypto/sha256"
        //"encoding/hex"
        "log"
        "sort"
        "strings"
)
const(
        //Pastry
        LEAFSLIMIT = 5
        RTLENGTH = 256
        PASTRYERROR_01 = "Leafset greater than default"
        PASTRYERROR_02 = "Peer already exist"
        PASTRYERROR_03 = "Peer doesn't exist"
        PASTRYERROR_04 = "No near peer found"


)

////////////////////////////////////////////////////////////////////////////////
////********************   PEER  *********************//////////////////////////
////////////////////////////////////////////////////////////////////////////////

type Peer struct{
        addr        string
        port        string

}
func (p *Peer) NewPeer(adr,prt string){
        fmt.Println(adr)
        fmt.Println(prt)
        p.addr = adr
        p.port = prt

}
///// Metodos GET

func (p *Peer) GetAddr() string{
        return p.addr
}

func (p *Peer) GetPort() string{
        return p.port
}

func (p *Peer) GetNumID() [32]byte{
        return hashme(p.addr)
}

func (p *Peer) GetID() string{
        return hexme(hashme(p.addr))
}



////////////////////////////////////////////////////////////////////////////////
////********************  PAsTRY PEER  *********************////////////////////
////////////////////////////////////////////////////////////////////////////////

type PPeer struct{
        *Peer
        parent          Peer
        children        map[string]Peer
        leafs           map[string]Peer
        rtcache         [256]map[string]Peer
}
///////////////Metodos

// Metodos Get

func (p *PPeer) GetParent() Peer{
        return p.parent
}

func (p *PPeer) GetChildrens() map[string]Peer{
        return p.children
}

func (p *PPeer) GetLeafs() map[string]Peer{
        return p.leafs
}

func (p *PPeer) GetRoutingTableCache() [256]map[string]Peer{
        return p.rtcache
}


// Metodos de comprobación

func (p *PPeer) IsMyLeaf(lkey string) bool{
        _,is := p.leafs[lkey]
        return is
}

func (p *PPeer) IsMyChild(ckey string) bool{
        _,is := p.children[ckey]
        return is
}

func (p *PPeer) IsInCache(pkey string) bool{
        dist := p.addressDistance(pkey)
        _,is := p.rtcache[dist][pkey]
        return is
}

// Metodos de configuración

func (p *PPeer) NewPPeer(adr, prt string, prnt Peer){
        log.Println("New Pastry Peer")
        p.NewPeer(adr,prt)
        p.parent = prnt

}


func (p *PPeer) NewLeafs( newLeafs map[string]Peer ){
        if len(newLeafs) <= LEAFSLIMIT {
                p.leafs = newLeafs
        } else {
                log.Println(PASTRYERROR_01)
        }

}

func (p *PPeer) AddChild( child Peer){
        if _,is := p.children[child.GetID()]; is == false {
                p.children[child.GetID()] = child
        } else {
                log.Println(PASTRYERROR_02)
        }
}

func (p *PPeer) RemoveChild( chkey string){
        if _,is := p.children[chkey]; is == true {
                delete(p.children, chkey)
        } else {
                log.Println(PASTRYERROR_03)
        }
}

func (p *PPeer) AddToCache( per Peer){
        if _,is := p.rtcache[per.GetID()]; is == false {
                p.rtcache[per.GetID()] = per
        } else {
                log.Println(PASTRYERROR_02)
        }
}

func (p *PPeer) RemoveFromCache( pkey string){
        dist := p.addressDistance(pkey)
        if _,is := p.rtcache[dist][pkey]; is == true {
                delete(p.rtcache[dist], pkey)
        } else {
                log.Println(PASTRYERROR_03)
        }
}

///// Métodos de Pastry


func (p *PPeer) SetRoutingTable() {

}

func (p *PPeer) SetLeafTable() {

}



func (p *PPeer) GetNearestPeer(toPeer string) Peer{
        if p.IsMyLeaf(toPeer){
                return p.leafs[toPeer]
        }
        dist := p.addressDistance(toPeer)
        nearSet := p.rtcache[dist]

        if peer:= p.nearPeerOfSet(toPeer, nearSet), peer != nil{
                return peer
        }
        log.Println(PASTRYERROR_04)
        return nil


}

func (p *PPeer) addressDistance(key string) int{

        dist := 0
        for i := range p.addr{
                if string(p.addr[i])==string(key[i]){
                        dist +=1
                } else{
                        break
                }
        }
        return dist

}

func (p *PPeer) nearPeerOfSet(toPeer string, peerSet map[string]Peer) Peer{
        ptmp := make([]string,0,len(peerSet))
        for key := range peerSet{
                append(ptmp,key)
        }
        sort.Strings(ptmp)
        for _,k := range ptmp{
                if strings.Compare(toPeer,k) != 1 {
                        return peerSet[k]
                }
        }
        return nil
}

//////////////////////////////////////////////////////////
///////////////////////Funciones Auxiliares //////////////


func tobin( value interface{} ) string {
        return fmt.Sprintf("%b",value)
}

func hexme( value interface{} ) string {
        return fmt.Sprintf("%x",value)

}

func hashme( value interface{} ) [32]byte{

        return sha256.Sum256([]byte(fmt.Sprintf("%v",value)))
}

func toString( value interface{} ) string{

        return fmt.Sprintf("%v",value)
}
