package workshop

import (
	"crypto/rand"
	"os"

	tj "github.com/Hyperledger-TWGC/tjfoc-gm/sm2"
	tjx509 "github.com/Hyperledger-TWGC/tjfoc-gm/x509"
)

type TJSM2 struct {
	PrivateKey *tj.PrivateKey
	PublicKey  *tj.PublicKey
}

func NewTJSM2() (*TJSM2, error) {
	PrivateKey, err := tj.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &TJSM2{PrivateKey: PrivateKey, PublicKey: &PrivateKey.PublicKey}, nil
}

func TJImportKey(privPEM []byte, pubPEM []byte) (*TJSM2, error) {
	PrivateKey, err := tjx509.ReadPrivateKeyFromPem(priv