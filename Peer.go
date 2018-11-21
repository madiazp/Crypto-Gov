/*
Definición de la clase peer, que solo caracteriza los atributos basicos del peer y metodos para modificarlos
y retornarlos
*/
package Peer
import (
        "fmt"
        "crypto/rsa"
)

type Peer struct {
  addr, id                     String
  knowns_peers, neighbors      map[String]String

}

// Métodos para fijar información

func (p *Peer) Set_Netwaddr(naddr String){
      p.addr = naddr
}

// Métodos de modificacion de atributos

func (p *Peer) Add_KnowPeer(id, naddr String){
  if knowns_peers != nil {
    knowns_peers[id] = naddr
  }
}

func (p *Peer) Add_Neighbor(id, naddr String){
  if neighbors != nil {
    neighbors[id] = naddr
  }
}

func (p *Peer) Del_KnowPeer(id String){
  delete(knowns_peers, id)
}

// Métodos para obtener info

func (p *Peer) Get_Id() String{
  return p.id
}

func (p *Peer) Get_addr() String{
  return p.addr
}

func (p *Peer) Get_KnowPeers() map[String]String {
  return p.knowns_peers
}

func (p *Peer) Get_Neighbors() map[String]String {
  return p.neighbors
}

/////////////////////////////////////////////////////////
