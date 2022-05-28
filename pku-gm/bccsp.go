package gmssl

import (
	. "github.com/Hyperledger-TWGC/pku-gm/gmssl"
	. "github.com/hyperledger/fabric/bccsp"
	. "hash"
	"strings"
)

// SM2PrivateKey
type SM2PrivateKey struct {
	*PrivateKey
	Password string
	skiHash  Hash
}

func (p *SM2PrivateKey) Bytes() ([]byte, error) {
	pem, err := p.GetPEM(SMS4,