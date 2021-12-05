package interop

import (
	"testing"
	"time"

	ccs "github.com/Hyperledger-TWGC/ccs-gm/sm3"
	pku "github.com/Hyperledger-TWGC/pku-gm/gmssl"
	tj "github.com/Hyperledger-TWGC/tjfoc-gm/sm3"
)

var base_format = "2006-01-02 15:04:05"

func TestSM3(t *testing.T) {
	// generate a rando