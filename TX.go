package TX

import (
        "fmt"
        "crypto/sha256"
)

type TX struct {

      content map[[]byte]Vote
      contHash []byte
      txPubKey []byte
      txSign []byte
      signAddr []byte
}
// funciones GET
func (t *TX) GetMsgs() map[[]byte]Vote {
      return t.content
}

func (t *TX) GetPub() []byte {
        return t.txPubKey
}

func (t *TX) GetSign() []byte{
        return t.txSign

}
func (t *TX) GetHash() []byte{
        return t.contHash
}

func (t *TX) GetAVote( h []byte) Vote {
        return t.content[h]
}

func (t *TX) GetAddr() []byte{
        return t.signAddr
}
//Funciones de Utilidad

func (t *TX) AddVote (v Vote) bool {
        h := v.GetAddr()
        if !t.Exist(h){
                t.content[h] = v
                return true
        }
       return false
}

func (t *TX) Exist (h []byte) bool {
        _,ok := t.content[h]
        return ok
}
func (t *TX) HashTheContent(){
        t.contHash = hashme(t.content)
}
func (t *TX) MakeAddr(){
        t.signAddr = hashme(t.txPubKey)
}
////////////////////////////////////////
//Vote


type Vote struct{
        value   string
        sign    []byte
        pubKey  []byte
        addr    []byte
}
func (v *Vote) NewVote(val string, sgn []byte ,pbky []byte) {
        v.value = val
        v.sign = sgn
        v.pubKey = pbky
        v.MakeAddr()

}
func (v *Vote) GetVal() string{
        return v.value
}

func (v *Vote) GetSign() []byte{
        return v.sign
}

func (v *Vote) GetPub() []byte{
        return v.pubkey
}

func (v *Vote) GetAddr() []byte{
        return v.addr
}
//metodos de utilidad

func (v *Vote) MakeAddr(){
        v.addr = hashme(v.pubKey)
}


///////////////////////////////////////////
/////////////Funciones auxiliares////////////

func hashme( value interface{} ) []byte{

        return sha256.Sum256([]byte(fmt.Sprintf("%v",value)))
}

func toString( value interface{} ) string{

        return fmt.Sprintf("%v",value)
}
