
package p2p
/*
Matias Díaz

Este package describe las estructura que se encargan de las conexiones entre pares
Node describe un nodo compuesto de un peer: su descriptor, handler: un auxiliar para las funciones de conexión
y un interruptor que para la rutina que escucha conexiones.

Handler describe a la clase que tiene las funciones para recibir y mandar conexiones entre pares, un nodo
lo usara como atributo auxiliar y mediante él hará todas las tareas de conexión
*/

import (
        "fmt"
        "net"
        "log"
        "io/ioutil"
        "time"
        "sync"

)
//Estructura del nodo
//peer tiene todos los datos que describen al nodo (el no es un peer)
//handler es un auxiliar que realiza las funciones de conexión
//interrupt es un flag(chanel) que si es utilizado le dice a la rutina de conexión
//que debe parar y cómo hacerlo según fue invocada la flag

type Node struct {
        peer Peer
        handler Handler
        interrupt chan int

}

// inicializa al nodo
func (n *Node) NewNode(adr,prt string ){
        n.peer.NewPeer(adr,prt)
        n.SetHandler()
        n.interrupt = make(chan int)
}
//Crea al handler correspondiente al nodo, se entrega a si mismo para iniciar al handler
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

func (n *Node) GetID() string{
        return n.peer.GetID()
}

//////////////// Metodos de Utilidad /////////////

// Inicia la rutina que escucha conexiones de entrada
// Dependiendo de la señal de interrupcion puede hacer Stop: Esperará a que todas
//las rutinas de conexión termine antes de terminar,
//Halt: terminará abruptamente
func (n *Node) Start(){
        var wg sync.WaitGroup
        wg.Add(1)
        go n.handler.Start(&wg)
        select{
        case a:= <-n.interrupt:
                if a == 1{
                        wg.Done()
                        log.Println("Stopped")
                        break
                } else{
                        log.Println("HALT")
                        return
                }
        }
        fmt.Println("Waiting?",wg)
        wg.Wait()

}
//Mandará una señal Stop a la rutina que escucha conexiones
func (n *Node) Stop(){
        log.Println("Stopping")
        n.interrupt <- 1
}
//Madará una señal Halt a la rutina que escucha conexiones
func (n *Node) Halt(){
        log.Println("Halting")
        n.interrupt <- 0
}
// Envía el mensaje msg al nodo tosend
func (n *Node) Send(msg string, tosend Node){
        n.handler.Send(msg,tosend)
}
//////////////////////////////////////////////////
//////////////////////////////////////////////////
/////////////////////////////////////////////////

//Estructura que describe al handler, se necesita la dirección IP y el puerto
//En que se escuchará conexiones entrantes del nodo que utilizará al handler
type Handler struct{

        addr            string
        port            string
}



//////////// Metodos /////////////////

//Inicializa al handler
func (h *Handler) NewHandler( n Node){
        h.addr = n.GetAddr()
        h.port = n.GetPort()




}
//////////Metodos de Utilidad /////////////

// crea un listener y se pone a escuchar en la dirección y puerto del nodo
// si una conexión entra la manda a un hilo que la procesará

func (h *Handler) Start(wg *sync.WaitGroup) {

        fmt.Println("Starting")
        fulladr := string(h.addr+":"+h.port)
        fmt.Println(fulladr)
        lst,err := net.Listen("tcp", fulladr)


        if err != nil{
                log.Println(err)
        }

        for{
                conn,err := lst.Accept()
                if err != nil {
                        log.Println("SOMETHING'S WRONG")
                        log.Println(err)
                        return
                }
                go h.HandleCon(conn, wg)
                wg.Add(1)

                fmt.Println("handled")

        }

        defer lst.Close()

}

// procesa el mensaje llegado
func (h *Handler) HandleCon(cn net.Conn, wg *sync.WaitGroup){
        log.Println("Received")

        buff,err := ioutil.ReadAll(cn)
        Fatal(err)
        fmt.Println(string(buff))
        wg.Done()

        defer cn.Close()

}
//Envia msg al nodo tonode
//Espera 3 segundos antes de fallar por time out
func (h *Handler) Send(msg string, toNode Node){

        wait, err := time.ParseDuration("3s")
        Fatal(err)
        conn, err := net.DialTimeout("tcp", toNode.GetAddr()+":"+toNode.GetPort(),wait)
        Fatal(err)
        buff := []byte(msg)
        fmt.Println("to send: %v",msg)
        _,err2 := conn.Write(buff)
        Fatal(err2)
        defer conn.Close()

}

///////////////Funciones auxiliares////////

func Fatal(er error){
        if er != nil{
                log.Fatal(er)
        }
}

func PError(er error){
        if er != nil{
                log.Println(er)
        }
}
