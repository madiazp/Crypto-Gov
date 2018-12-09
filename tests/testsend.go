package main

import(

        p2p "../p2p"

)

func main(){

        var nod p2p.Node
        var nod2 p2p.Node
        nod.NewNode("127.0.0.1","9092")
        nod2.NewNode("127.0.0.1","9091")

        nod.Start()
        nod.Send("oli",nod2)



}
