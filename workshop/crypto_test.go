package workshop_test

import (
	"github.com/Hyperledger-TWGC/fabric-gm-plugins/workshop"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var source, priv, pub workshop.SM2
var SM4Key  workshop.SM4
var priFile = "priv.pem"
var pubFile = "pub.pem"
var msg = []byte("2021-