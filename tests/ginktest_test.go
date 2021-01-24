package tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"ginktest/tests"
)

var _ = Describe("Ginktest", func() {
	BeforeEach(func() {
		hello()
	})

	Describe("An example test", func() {
		It("might pass", func(){
			Expect(true==true)
		})
	})
})
