package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var patcher string

var _ = BeforeSuite(func() {
	var err error
	patcher, err = gexec.Build("github.com/pivotal-cf-experimental/knit")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestApplyPatches(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApplyPatches Suite")
}
