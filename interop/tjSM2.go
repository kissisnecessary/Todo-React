package interop

import (
	"crypto/rand"

	tj "github.com/Hyperledger-TWGC/tjfoc-gm/sm2"
	tjx509 "github.com/Hyperledger-TWGC/tjfoc-gm/x509"
)

type TJSM2 struct {
	PrivateKey *tj.PrivateKey
	PublicKey  *tj.PublicKey
}

func NewTJSM2() 