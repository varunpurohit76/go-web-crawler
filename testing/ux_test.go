package testing

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/varunpurohit76/crawler/base"
	"github.com/varunpurohit76/crawler/graph"
	"github.com/varunpurohit76/crawler/scrapper"
	"github.com/varunpurohit76/crawler/sitemap"
)

func TestSiteMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "sitemap")
}

var _ = BeforeSuite(func() {
	TestServicesInit(scrapper.MockPageUrlExtract, graph.JsonView)
})

var _ = Describe("sitemap as a service", func() {
	Context("basic ux MockPageUrlExtract JsonView", func() {
		It("www.monzo.com depth=0", func() {
			rc := base.NewRequestContext()
			req := &sitemap.SiteMapRequest{
				Url:   "https://www.monzo.com",
				Depth: 0,
			}
			res := sitemap.Service.Build(rc, req.Url, req.Depth)
			resObj := res.(*graph.SiteMapUrlNodeView)
			Expect(resObj.Link).To(BeEquivalentTo("https://www.monzo.com"))
			Expect(resObj.Children).To(BeNil())
		})
		It("www.monzo.com depth=1", func() {
			rc := base.NewRequestContext()
			req := &sitemap.SiteMapRequest{
				Url:   "https://www.monzo.com",
				Depth: 1,
			}
			type void struct{}
			var member void
			res := sitemap.Service.Build(rc, req.Url, req.Depth)
			resObj := res.(*graph.SiteMapUrlNodeView)
			Expect(resObj.Link).To(BeEquivalentTo("https://www.monzo.com"))
			Expect(len(resObj.Children)).To(BeEquivalentTo(2))
			chLevel1 := resObj.Children
			chLevel1ValidUrl := make(map[string]void)
			chLevel1ValidUrl["http://www.monzo.com/a"] = member
			chLevel1ValidUrl["http://www.monzo.com/b"] = member
			for _, c := range chLevel1 {
				Expect(chLevel1ValidUrl[c.Link]).To(BeEquivalentTo(member))
				Expect(c.Children).To(BeNil())
			}
		})
		It("www.monzo.com depth=2", func() {
			rc := base.NewRequestContext()
			req := &sitemap.SiteMapRequest{
				Url:   "https://www.monzo.com",
				Depth: 2,
			}
			type void struct{}
			var member void
			res := sitemap.Service.Build(rc, req.Url, req.Depth)
			resObj := res.(*graph.SiteMapUrlNodeView)
			Expect(resObj.Link).To(BeEquivalentTo("https://www.monzo.com"))
			Expect(len(resObj.Children)).To(BeEquivalentTo(2))
			chLevel1 := resObj.Children
			var chLevel2 []*graph.SiteMapUrlNodeView
			chLevel1ValidUrl := make(map[string]void)
			chLevel1ValidUrl["http://www.monzo.com/a"] = member
			chLevel1ValidUrl["http://www.monzo.com/b"] = member
			for _, c := range chLevel1 {
				Expect(chLevel1ValidUrl[c.Link]).To(BeEquivalentTo(member))
				chLevel2 = append(chLevel2, c.Children...)
			}
			chLevel2ValidUrl := make(map[string]void)
			chLevel2ValidUrl["http://www.monzo.com/a1"] = member
			chLevel2ValidUrl["http://www.monzo.com/a2"] = member
			chLevel2ValidUrl["http://www.monzo.com/b1"] = member
			chLevel2ValidUrl["http://www.monzo.com/a2"] = member
			for _, c := range chLevel2 {
				Expect(chLevel2ValidUrl[c.Link]).To(BeEquivalentTo(member))
				Expect(c.Children).To(BeNil())
			}
		})
		It("www.monzo.com depth=3", func() {
			rc := base.NewRequestContext()
			req := &sitemap.SiteMapRequest{
				Url:   "https://www.monzo.com",
				Depth: 3,
			}
			type void struct{}
			var member void
			res := sitemap.Service.Build(rc, req.Url, req.Depth)
			resObj := res.(*graph.SiteMapUrlNodeView)
			Expect(resObj.Link).To(BeEquivalentTo("https://www.monzo.com"))
			Expect(len(resObj.Children)).To(BeEquivalentTo(2))
			chLevel1 := resObj.Children
			var chLevel2 []*graph.SiteMapUrlNodeView
			var chLevel3 []*graph.SiteMapUrlNodeView
			chLevel1ValidUrl := make(map[string]void)
			chLevel1ValidUrl["http://www.monzo.com/a"] = member
			chLevel1ValidUrl["http://www.monzo.com/b"] = member
			for _, c := range chLevel1 {
				Expect(chLevel1ValidUrl[c.Link]).To(BeEquivalentTo(member))
				chLevel2 = append(chLevel2, c.Children...)
			}
			chLevel2ValidUrl := make(map[string]void)
			chLevel2ValidUrl["http://www.monzo.com/a1"] = member
			chLevel2ValidUrl["http://www.monzo.com/a2"] = member
			chLevel2ValidUrl["http://www.monzo.com/b1"] = member
			chLevel2ValidUrl["http://www.monzo.com/a2"] = member
			for _, c := range chLevel2 {
				Expect(chLevel2ValidUrl[c.Link]).To(BeEquivalentTo(member))
				chLevel3 = append(chLevel3, c.Children...)
			}
			chLevel3ValidUrl := make(map[string]void)
			chLevel3ValidUrl["http://www.monzo.com/a11"] = member
			chLevel3ValidUrl["http://www.monzo.com/a12"] = member
			chLevel3ValidUrl["http://www.monzo.com/a21"] = member
			chLevel3ValidUrl["http://www.monzo.com/a22"] = member
			chLevel3ValidUrl["http://www.monzo.com/b11"] = member
			chLevel3ValidUrl["http://www.monzo.com/b12"] = member
			chLevel3ValidUrl["http://www.monzo.com/a21"] = member
			chLevel3ValidUrl["http://www.monzo.com/a22"] = member
			for _, c := range chLevel3 {
				Expect(chLevel2ValidUrl[c.Link]).To(BeEquivalentTo(member))
				Expect(c.Children).To(BeNil())
			}
		})
	})
})
