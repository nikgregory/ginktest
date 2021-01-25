package tests_test

import (
	ginktests "ginktest/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Ginktest", func() {
	BeforeEach(func() {
		ginktests.Hello()
	})

	Describe("An example test", func() {
		It("might pass", func(){
			Expect(true==true)
		})
	})
})
