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

	BeforeSuite(func() {
		tmpDir, err = ioutil.TempDir("", "workshop-e2e-")
		Expect(err).NotTo(HaveOccurred())

		serverBin, err = gexec.Build("./server")
		Expect(err).NotTo(HaveOccurred())

		clientBin, err = gexec.Build("./client")
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeEach(func() {
		cmd := exec.Command(clientBin, tmpDir, "generate")
		clientSession, err = gexec.Start(cmd, nil, nil)
		Eventually(clientSession.Out).Should(Say("generate key pair at " + tmpDir))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if serve