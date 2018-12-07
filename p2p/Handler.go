package Handler

import (
        "fmt"
        "crypto/sha256"
        "net"
        "log"
        p2p "."
)

type Handler struct{

        mynode          p2p.Node
        interrupts      chan <- bool
}


//////////// Metodos /////////////////
func (h *Hanlder) NewHandler( n p2p.Node){
        if h.mynode == nil {
                h.mynode = n
        }
}

func (h *Handler) listen(lst Listener ){

        conn,err := lst.Accept()
        if err != nil {
                log.Print(err)
                continue

        go h.HandleCon(conn)

}

func (h *Handler) Start() {
        h.interrupts <- false
        list,err := net.Listener("tcp", mynode.addr+":"+mynode.port)
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


func (h *Handler) HandleCon(cn Conn){
        defer cn.Close()
        buff := ioutil.ReadAll(cn)
        fmt.Sprintf("%v",buff)

}
