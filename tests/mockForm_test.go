package tests_test

import (
	// . "mockForm/app"
	// "mockForm/app/controllers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
 . "github.com/sclevine/agouti/matchers"
)

var _ = Describe("MockForm", func() {
	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("Should show the page", func(){
		Expect(page.Navigate("http://localhost:9000")).To(Succeed())
		Expect(page).To(HaveURL("http://localhost:9000/"))

	})
	It("Should have a heading", func(){
		Expect(page.Navigate("http://localhost:9000")).To(Succeed())
		Eventually(page.Find(".header")).Should(HaveText("API Deployment Instructions"))
	})
	It("Should have a submit button and route to a new page", func(){
		Expect(page.Navigate("http://localhost:9000")).To(Succeed())
		Expect(page.Find(".submitBtn").Click()).To(Succeed())
		Expect(page).To(HaveURL("http://localhost:9000/verify"))
	})

})
