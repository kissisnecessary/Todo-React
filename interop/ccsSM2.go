package interop

import (
	"crypto/rand"

	ccs "github.com/Hyperledger-TWGC/ccs-gm/sm2"
	ccsutils "github.com/Hyperledger-TWGC/ccs-gm/utils"
)

type CCSSM2 struct {
	PrivateKey *ccs.Pri