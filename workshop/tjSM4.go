package workshop

import tj "github.com/Hyperledger-TWGC/tjfoc-gm/sm4"

type TJSM4 struct {
	Key []byte
}

func NewTJSM4() (*TJSM4, error) {
	key := []byte("0123456789abcdef")
	return &TJSM4{Key: key}, nil
}
func (instance *TJSM4) Encrypt