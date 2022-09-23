package workshop

import (
	"crypto/rand"

	ccs "github.com/Hyperledger-TWGC/ccs-gm/sm2"
	ccsutils "github.com/Hyperledger-TWGC/ccs-gm/utils"
)

type CCSSM2 struct {
	PrivateKey *ccs.PrivateKey
	PublicKey  *ccs.PublicKey
}

func NewCCSSM2() (*CCSSM2, error) {
	PrivateKey, err := ccs.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &CCSSM2{PrivateKey: PrivateKey, PublicKey: &PrivateKey.PublicKey}, nil
}

func CCSImportKey(privPEM []byte, pubPEM []byte) (*CCSSM2, error) {
	Priv