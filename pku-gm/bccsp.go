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
	pem, err := p.GetPEM(SMS4, p.Password)
	return []byte(pem), err
}
func (p *SM2PrivateKey) Symmetric() bool {
	return false
}
func (p *SM2PrivateKey) Private() bool {
	return true
}

func (p *SM2PrivateKey) PublicKey() (Key, error) {
	publicKey, err := p.GetPublicKey()
	if err != nil {
		return nil, err
	}

	return &SM2PublicKey{Key: publicKey}, nil
}
func (p *SM2PrivateKey) SKI() []byte {
	text, err := p.Ge