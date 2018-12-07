package Node

import (
        "fmt"
        "net"
        "log"
        "io/ioutil"
        "time"

        p2p ".."
)
type Node struct {
        peer p2p.Peer
        handler Handler
}


func (n *Node) NewNode(adr,prt string, ){
        n.peer.NewPeer(adr,prt)
}

func (n *Node) SetHandler(){
        if n.peer.GetAddr() != ""{
                n.handler.NewHandler(*n)
        } else {
                log.Print("no address")
        }

}
///////////Metodos GET ////////////

func (n *Node) GetAddr() string{
        return n.peer.GetAddr()
}

func (n *Node) GetPort() string{
        return n.peer.GetPort()
}
func (n *Node) GetID()[32]byte{
        return n.peer.GetID()
}
//////////////// Metodos de Utilidad /////////////

func (n *Node) Start(){

        n.handler.Start()
}

//////////////////////////////////////////////////
//////////////////////////////////////////////////
/////////////////////////////////////////////////

type Handler struct{

        addr            string
        port            string
        interrupts      chan  bool
}


//////////// Metodos /////////////////
func (h *Handler) NewHandler( n Node){

        h.addr = n.GetAddr()
        h.port = n.GetPort()

}
//////////Metodos de Utilidad ///////////77

func (h *Handler) listen(lst net.Listener ){

        conn,err := lst.Accept()
        if err != nil {
                log.Print(err)
                return
        }

        go h.HandleCon(conn)

}

func (h *Handler) Start() {
        h.interrupts <- false
        fulladr := string(h.addr+":"+h.port)
        list,err := net.Listen("tcp", fulladr)
        if err != nil{
                log.Fatal(err)
        }
        for{
                go h.listen(list)
                if <-h.interrupts{
                        break
                }
        }


}

func (h *Handler) Stop(){
        h.interrupts <- true
}


func (h *Handler) HandleCon(cn net.Conn){
        defer cn.Close()
        buff,err := ioutil.ReadAll(cn)
        if err != nil{
                log.Fatal(err)
        }
        fmt.Sprintf("%v",buff)

}

func (h *Handler) Send(msg string){

        wait, err := time.ParseDuration("3s")
        conn, err := net.DialTimeout("tcp", h.addr+":"+h.port,wait)
        buff := []byte(msg)
        _,err = conn.Write(buff)
        if err != nil{
                log.Fatal(err)
        }

}
