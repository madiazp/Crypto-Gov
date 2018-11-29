package block

import (
        "fmt"
        "crypto/sha256"
        "crypto/rsa"
)

type TX struct {
        data    []string
}

func (t *TX) GetData() []string{
        return t.data
}

func (t *TX) AddDara(msg string) {
        append(t.data, msg)
}

func (t *TX) VerifyData() bool{
        for _, msg := range t.data[:]{
                if data != "True"{
                        return false
                }
        }
        return true
}

type FirmedTX struct {
        tx      TX
        index   []byte
        sign    []byte
        pubkey  rsa.PublicKey


}
