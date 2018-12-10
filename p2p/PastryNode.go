package p2p

import(
        "fmt"
)

const (
        // MENSAJES

        // JOIN & NETWORK
        JOIN =  "00"
        LEAFS = "01"
        WHOIS = "02"
        ITIS  = "03"

)
type PNode struct{
        nod             Node
        parent          Peer
        children        map[string]Peer
        leafs           map[string]Peer
        rtcache         map[string]Peer
}
////////// Metodos Get

func (p *PNode) GetNode() Node{
        return p.nod
}

func (p *PNode) GetLeafs() map[string]Peer{
        return p.leafs
}

func (p *PNode) GetRoutingTableCache() map[string]Peer{
        return p.rtchache
}

/////////// Metodos de configuracion de la red

func (p *PNode) NewPNode(nd Node, prnt Peer){

        p.nod = nd
        p.nod.SetPNode(p)
        p.parent = prnt
        p.AskToJoin()


}

func (p *PNode) AskToJoin(){
        log.Println("Asking to join to "+p.parent.GetID())
        p.send(JOIN+p.nod.GetID()+p.nod.GetPort(), p.parent)

}

func (p *PNode) Start(){
        log.Println("Starting")
        var wg sync.WaitGroup
        go p.nod.Start(&wg,p)
        log.Println("Finished")

}
///// Metodos de envio

func (p *PNode) send(code string, pr Peer){
        p.nod.Send(code,pr)
}

//// Handler
func (p *PNode) HandleCon(cn net.Conn, wg *sync.WaitGroup){
        log.Println("Received")

        buff,err := ioutil.ReadAll(cn)
        Fatal(err)
        fmt.Println(string(buff))
        wg.Done()

        defer cn.Close()

}
