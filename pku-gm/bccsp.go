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
	text, err := p.GetText()
	PanicError(err)
	p.skiHash.Reset()
	p.skiHash.Write([]byte(text))
	sum := p.skiHash.Sum(nil)
	p.skiHash.Reset()
	return sum
}

// SM2PublicKey
type SM2PublicKey struct {
	Key     *PublicKey
	skiHash Hash
}

func (p *SM2PublicKey) Bytes() ([]byte, error) {
	pem, err := p.Key.GetPEM()
	return []byte(pem), err
}
func (p *SM2PublicKey) SKI() []byte {

	text, err := p.Key.GetText()
	PanicError(err)
	p.skiHash.Reset()
	p.skiHash.Write([]byte(text))
	sum := p.skiHash.Sum(nil)
	p.skiHash.Reset()
	return sum
}

func (p *SM2PublicKey) Symmetric() bool {
	return false
}
func (p *SM2PublicKey) Private() bool {
	return false
}
func (p *SM2PublicKey) PublicKey() (Key, error) {
	return p, nil
}

// Software-based GM Suite
type GMSWSuite struct {
	KeyStore
}

// KeyGen generates a Key using opts. FIXME logic correct?
func (s *GMSWSuite) KeyGen(opts KeyGenOpts) (k Key, err error) {
	var algorithm = opts.Algorithm()

	// fall to default
	if algorithm == "" {
		algorithm = "sm2p256v1"
	}
	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", algorithm},
		{"ec_param_enc", "named_curve"},
	}
	// TODO factory to support multiple Key type
	sm2sk, err := GeneratePrivateKey("EC", sm2keygenargs, nil)
	if !opts.Ephemeral() {
		// Store the Key
		err = s.StoreKey(k)
		if err != nil {
			return nil, err
		}
	}
	sm3 := New()
	return &SM2PrivateKey{PrivateKey: sm2sk, skiHash: sm3}, nil
}

// KeyDeriv derives a Key from k using opts.
// The opts argument should be appropriate for the primitive used.
func (s *GMSWSuite) KeyDeriv(k Key, opts KeyDerivOpts) (dk Key, err error) {
	panic("To be Implement") // TODO
}

// KeyImport imports a Key from its raw representation using opts.
// The opts argument should be appropriate for the primitive used