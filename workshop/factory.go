
package workshop

import (
	"io/ioutil"

	"github.com/Hyperledger-TWGC/tjfoc-gm/x509"
)

var TJ = "TJ"
var PKU = "PKU"
var CCS = "CCS"

func GenerateSM2Instance(sourceDef string) (SM2, error) {
	if sourceDef == TJ {
		return NewTJSM2()
	}
	if sourceDef == CCS {