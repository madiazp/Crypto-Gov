package block

import (
        "fmt"
        "crypto/sha256"
)
///// definicion del bloque
type Ublock struct {

      prevHash []byte
      merkelTree []byte
      timeStamp string
      payload TX

}

func (u *Ublock) GetMSGS() []string{
  return payload.GetMsgs()
}

//////////////////////////////////

type TX struct {

      msgs []string
}

func (t *TX) GetMsgs() []string {
      return t.msgs
}

func (t *TX) AddMsg (mg string) {
      append(t.msgs,mg)
}

func (t *TX) VerifyMSGS() (bool)  {
      for i:= range t.msgs {
        if t.msgs[i] != "True" {
          return false
        }
      }
      return true
}
/////////////////////////////////////

type Block struct {
      ublk Ublock
      selfhash []byte
}

func (b *Block) GetMsgs() []String {
      return b.ublk.GetMSGS()
}

func (b *Block) VerifyTXS() (bool, error)  {
    if !


}

func (b *Block) SelfSign() error {
      var err error
      b.selfhash, err = sha256.Sum256(fmt.Sprintf("%v",b.ublk))
      return err
}

func (b *Block) SelfVerifySign() (bool, error) {
      test, err := sha256.Sum256(fmt.Sprintf("%v",b.ublk))
      if err != nil || test != b.selfhash{
        return false, err
      }
      return true
}
func (b *Block) SelfVerify() (bool, error) {
      if (b.SelfVerifySign()) { return false,nil }



}

////////////////////////////////////////////7
