package sriov

import (
	"github.com/MangoBoost/sriov-cni/pkg/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sriov Suite")
}

var _ = BeforeSuite(func() {
	// create test sys tree
	err := utils.CreateTmpSysFs()
	Expect(err).Should(Succeed())
})

var _ = AfterSuite(func() {
	err := utils.RemoveTmpSysFs()
	Expect(err).Should(Succeed())
})
