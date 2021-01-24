package ginktest_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinktest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginktest Suite")
}
