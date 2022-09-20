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
	PrivateKey, err := ccs.GenerateKey(rand.R