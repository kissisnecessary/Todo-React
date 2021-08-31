
package interop

import (
	pku "github.com/Hyperledger-TWGC/pku-gm/gmssl"
)

type PKUSM2 struct {
	PrivateKey *pku.PrivateKey
	PublicKey  *pku.PublicKey
}

func NewPKUSM2() (*PKUSM2, error) {
	/* SM2 key pair operations */
	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	sm2sk, err := pku.GeneratePrivateKey("EC", sm2keygenargs, nil)
	if err != nil {
		return nil, err
	}
	sm2pkpem, err := sm2sk.GetPublicKeyPEM()
	if err != nil {
		return nil, err
	}
	sm2pk, err := pku.NewPublicKeyFromPEM(sm2pkpem)
	return &PKUSM2{PrivateKey: sm2sk, PublicKey: sm2pk}, nil
}

func PKUImport(privPEM []byte, pubPEM []byte) (*PKUSM2, error) {
	sm2sk, err := pku.NewPrivateKeyFromPEM(string(privPEM), "")
	if err != nil {
		return nil, err
	}
	sm2pk, err := pku.NewPublicKeyFromPEM(string(pubPEM))
	return &PKUSM2{PrivateKey: sm2sk, PublicKey: sm2pk}, nil
}

func (instance *PKUSM2) ExportKey() (privPEM []byte, pubPEM []byte, err error) {
	privStr, err := instance.PrivateKey.GetPEM("", "")
	if err != nil {
		return
	}
	pubStr, err := instance.PrivateKey.GetPublicKeyPEM()
	privPEM = []byte(privStr)
	pubPEM = []byte(pubStr)
	return
}

func (instance *PKUSM2) Encrypt(msg []byte) ([]byte, error) {
	sm2ciphertext, err := instance.PublicKey.Encrypt("sm2encrypt-with-sm3", []byte(msg), nil)
	return sm2ciphertext, err