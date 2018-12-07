package Node

import (
        "fmt"
        "crypto/sha256"
        pr "Peer"
)

type Node struct{
        pr.Peer
        handler Handler
}
