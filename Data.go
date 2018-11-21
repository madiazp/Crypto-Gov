// Estructuras de datos basicas para almacenar informacion en los nodos revisores
// Por el momento se esta fijando el algoritmo de hash a sha256 (El cual puede o no cambiar)
// y el metodo de firma digital a rsa.

package Data
import (
        "fmt"
        "crypto/rsa"
        "crypto"
)
// Estructura de la transacción, por ahora lleva un mensaje en texto plano, la idea es que
// lleve un contrato inteligente a ejecutar
type TX struct {

  value  String

}

////////////////////////////////////////////////////////


//Msg es la estructura de los mensajes
// contiene el valor del mensaje, que es la transacción (El payload), y el conjunto de firmas y pubkeys de
// los revisores que han confirmado el mensajes.
// Se debe entender que las firmas están en cadena (y por lo tanto en orden) de la forma:
//
//         [Sn= sign(H(Sn-1),PKn), Sn-1= sign(H(Sn-2),PKn-1),..., S1=sign(H(Value),PK1)]
//
// Donde H() es el metodo de Hash (SHA256 por ahora) y PKn es la private Key del validador n
// Notar que para validar estas firmas solo se necesita el valor "Value" inicial (El payload) y las llaves públicas
// de los validadores



type Msg struct {

  val        TX            // El payload
  signatures [][]byte      // el slice de firmas de los validadores
  pubkeys    []rsa.PubKeys // el slice con las llaves publicas de los validadores

}
// verifica todas las firmas en cadena del mensaje, retorna false si una de las firmas no concuerda
// returna true ssi todas las firmas concuerdan, notar que la cantidad de firmas debe ser igual  a la cantidad de pubkeys
func (e *Msg) MsgVerify() (bool, err){

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
// metodo que agrega la firma del revisor al mensaje, esta firma se hace respecto a la ultima firma contenida
//en el mensaje, en el caso de ser el primer revisor se usa al valor del mensaje
  func (e *Msg) SignValue( pub rsa.PublicKey, prv rsa.PrivateKey) error {
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

// funcion que setea el value
// notar que si se cambia el valor de una entrada previamente firmada por un validador, la primera firma
// será inválida, lo que invalidará toda el mensaje

func (e *Msg) SetValue( val TX){
    e.value = val
  }
func (e *Msg) getVal() TX{
  return e.val
}
func (e *Msg)getId() []byte{
  return sha256.Sum256([]byte(e.val))
}
func (e *Msg) getDepth() int {
  return len(e.signatures)
  }

//////////////////////////////////////////////////////////////////////////

// Estructura que guarda el cache de transacciones en el nodo
// simplifica la tarea de validar, buscar la transacción y actualizar las validaciones
// Este cache NO reemplaza la suma total

type TXEntry struct {
  msge Msg
  id := msge.getId()
  nValidators :=  msge.getDepth()
}
// updatea las firmas
func (t *TXEntry) UpdateMsg(msg Msg){
  if msg.getId() == t.id {
    t.msge = msg
    t.nValidators = msg.getDepth()
  }
}
