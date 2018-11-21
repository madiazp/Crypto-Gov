package Data
import (
        "fmt"
        "crypto/rsa"
        "crypto"
)

type TX struct {

  value  String

}

type Entry struct {

  val        TX
  signatures [][]byte
  pubkeys    []rsa.PubKeys

}

func (e *Entry) EntryVerify() (bool, err){

  if len(e.signatures) != len(e.pubkeys){
    return false
  }
  var pub rsa.PublicKey
  var sign []byte
  var err error
  hashed := sha256.Sum256([]byte(e.TX))

  for i := range e.signatures {
    pub = e.pubkeys[i]
    sign = e.signatures[i]
    err = rsa.VerifyPKCS1v15( &pub, cypto.SHA256, hashed[:], sign)
    if err != nil {
      return false
    }
    hashed = sha256.Sum256(sign)
    }
    return true
  }

  func (e *Entry) SignValue( pub rsa.PublicKey, prv rsa.PrivateKey) error {
    var hashed []byte
    if e.signatures != nil {
      hashed = sha256.Sum256(e.signatures[len(e.signatures)-1])
    }  else {
      hashed = sha256.Sum256([]byte(e.TX))
    }
    sign, err := rsa.SignPKCS1v15(rand.Reader, prv, crypto.SHA256, hashed[:])
  }
  if err != nil {
    return
  }
  append(e.signatures,sign)
  append(e.pubkeys,pub)


}
