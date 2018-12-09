package main

import(
        "fmt"
        p2p "../p2p"
        "time"


)

func main(){

        var nod p2p.Node
        fmt.Println("inicio")
        nod.NewNode("127.0.0.1","9091")


        go nod.Start()
        time.Sleep(3 * time.Second)
        nod.Stop()
        time.Sleep(1 * time.Second)
        fmt.Println("fin")
}
