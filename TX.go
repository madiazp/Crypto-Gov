package TX

import (
        "fmt"
        "crypto"
        "crypto/sha256"
        "crypto/rsa"
        "crypto/rand"

)

type TX struct {

        content map[string]Vote
        contHash [32]byte
        txPubKey rsa.PublicKey
        txSign []byte
        signAddr [32]byte
}

func (t *TX) NewTX( cnt map[string]Vote, pbky rsa.PublicKey, sgn []byte ){
        t.content = cnt
        t.txPubKey = pbky
        t.txSign = sgn
        t.HashTheContent()
        t.MakeAddr()
}
// funciones GET

func (t *TX) GetMsgs() map[string]Vote {
        return t.content
}

func (t *TX) GetPub() rsa.PublicKey {
        return t.txPubKey
}

func (t *TX) GetSign() []byte{
        return t.txSign

}
func (t *TX) GetHash() [32]byte{
        return t.contHash
}

func (t *TX) GetAVote( h string) Vote {
        return t.content[h]
}

func (t *TX) GetAddr() [32]byte{
        return t.signAddr
}
//Funciones de Utilidad


func (t *TX) Exist (hs string) bool {
        _,ok := t.content[hs]
        return ok
}

func (t *TX) HashTheContent(){
        t.contHash = hashme(t.content)
}

func (t *TX) MakeAddr(){
        t.signAddr = hashme(t.txPubKey)
}

func (t *TX) MakeTX(cnt map[string]Vote, prky rsa.PrivateKey, pbky rsa.PublicKey) error {

        tosign := hashme(cnt)
        rng := rand.Reader
        sign,err := rsa.SignPKCS1v15(rng, &prky, crypto.SHA256, tosign[:])
        if err != nil{
                return err
        }
        t.NewTX(cnt, pbky, sign)
        return nil
}
func (t *TX) ExecTX() bool {

        toverify := hashme(t.content)
        err := rsa.VerifyPKCS1v15( &t.txPubKey, crypto.SHA256, toverify[:], t.txSign)
        if err != nil {
                return false
        }
        return true
}

func (t *TX) TXSize() int{
        return len(t.content)
}
////////////////////////////////////////
//Vote


type Vote struct{

        value   string
        sign    []byte
        pubKey  rsa.PublicKey
        addr    [32]byte

}

func (v *Vote) NewVote(val string, sgn []byte ,pbky rsa.PublicKey) {

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

func (v *Vote) GetPub() rsa.PublicKey{
        return v.pubKey
}

func (v *Vote) GetAddr() [32]byte{
        return v.addr
}
//metodos de utilidad

func (v *Vote) MakeAddr(){
        v.addr = hashme(v.pubKey)
}

func (v *Vote) MakeVote( val string, prky rsa.PrivateKey, pbky rsa.PublicKey) error {

        tosign := hashme(val)
        rng := rand.Reader
        sign,err := rsa.SignPKCS1v15(rng, &prky, crypto.SHA256, tosign[:])
        if err != nil {
                return err
        }
        v.NewVote(val, sign, pbky)
        return nil
}

func (v *Vote) ExecVote() bool {

        toverify := hashme(v.value)
        err := rsa.VerifyPKCS1v15( &v.pubKey, crypto.SHA256, toverify[:], v.sign)
        if err != nil {
                return false
        }
        return true
}
///////////////////////////////////////////
/////////////Funciones auxiliares////////////

func hashme( value interface{} ) [32]byte {

        return sha256.Sum256([]byte(fmt.Sprintf("%v",value)))
}

func toString( value interface{} ) string{

        return fmt.Sprintf("%v",value)
}
