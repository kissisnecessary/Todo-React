package workshop_test

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var (
	PrivFile, PubFile            *os.File
	tmpDir, serverBin, clientBin string
	serverSession, clientSession *gexec.Session
)

var _ = Describe("Server", func() {

	BeforeSuite(