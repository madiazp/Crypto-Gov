
package p2p
/*
Matias Díaz

Este package describe las estructura que se encargan de las conexiones entre pares
Node describe un nodo compuesto de un peer: su descriptor, handler: un auxiliar para las funciones de conexión
y un interruptor que para la rutina que escucha conexiones.

Node describe a la clase que tiene las funciones para recibir y mandar conexiones entre pares, un nodo
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
const (
        // MENSAJES

        // JOIN & NETWORK
        JOIN =  "00"
        LEAFS = "01"
        WHOIS = "02"
        ITIS  = "03"

)
//Estructura del nodo
//peer tiene todos los datos que describen al nodo (el no es un peer)
//handler es un auxiliar que realiza las funciones de conexión
//interrupt es un flag(chanel) que si es utilizado le dice a la rutina de conexión
//que debe parar y cómo hacerlo según fue invocada la flag

type Node struct {
        peer PPeer
        interrupt chan int

}

// inicializa al nodo
func (n *Node) NewNode(adr,prt string ){

        log.Println("Adding node")
        log.Println(adr+":"+prt)
        n.peer.NewPeer(adr,prt)
        n.interrupt = make(chan int)


}
//Crea al handler correspondiente al nodo, se entrega a si mismo para iniciar al handler

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
// las rutinas de conexión termine antes de terminar,
// Halt: terminará abruptamente
// Debe ser invocado con un waitgroup seteado para esperar su término
func (n *Node) Start(wg *sync.WaitGroup){

        wg.Add(1)
        go n.Listen(wg)
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

// crea un listener y se pone a escuchar en la dirección y puerto del nodo
// si una conexión entra la manda a un hilo que la procesará

func (n *Node) Listen(wg *sync.WaitGroup) {

        fmt.Println("Starting")
        fulladr := string(n.GetAddr()+":"+n.GetPort())
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
                go n.HandleCon(conn, wg)
                wg.Add(1)

                fmt.Println("handled")

        }

        defer lst.Close()

}

// procesa el mensaje llegado

//Envia msg al nodo tonode
//Espera 3 segundos antes de fallar por time out
func (n *Node) Send(msg string, toNode PPeer){

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

func (n *Node) HandleCon(cn net.Conn, wg *sync.WaitGroup){
        log.Println("Received")

        buff,err := ioutil.ReadAll(cn)
        Fatal(err)
        fmt.Println(string(buff))
        wg.Done()

        defer cn.Close()

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
