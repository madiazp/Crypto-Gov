package main

import (
	"fmt"

	p2p "../p2p"
	//"time"
)

func main() {
	nod := new(p2p.PPeer)
	nod.NewPPeer("127.0.0.1", "9091")

	/*nod := p2p.PPeer{
	          p2p.Peer{
	                  addr:   "127.0.0.1",
	                  port:   "9091",
	                  id :    "",

	          },
	          parent:         nil,
	          children:       make(map[string]p2p.Peer),
	          leafs:          make(map[string]p2p.Peer),
	          rtcache:        make(map[string]p2p.Peer),
	  }
	*/
	fmt.Println("inicio")
	fmt.Println(nod.GetID())

}
