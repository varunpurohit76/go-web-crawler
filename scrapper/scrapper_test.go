package scrapper

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/varunpurohit76/crawler/base"
)



func TestScrapper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "scrapper")
}

var _ = Describe("scrapper", func() {
	var scrapper Scrapper
	rc := base.NewRequestContext()

	Context("PageUrlExtract", func() {
		It("should return links", func() {
			scrapper = ScrapperFactory(PageUrlExtract)
			links, err := scrapper.Scrap(rc, "https://www.monzo.com")
			Expect(err).To(BeNil())
			Expect(links).NotTo(BeNil())
		})
	})

	Context("MockPageUrlExtract", func() {
		It("should return mock links", func() {
			scrapper = ScrapperFactory(MockPageUrlExtract)
			links, err := scrapper.Scrap(rc, "https://www.monzo.com")
			Expect(err).To(BeNil())
			Expect(links).To(BeEquivalentTo([]string{"https://www.monzo.com/a", "https://www.monzo.com/b"}))
		})
	})
})
