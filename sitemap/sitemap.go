package sitemap

import (
	"github.com/varunpurohit76/crawler/base"
	. "github.com/varunpurohit76/crawler/graph"
)

var graph Graph
var viewType int
var Service = new(SiteMapImpl)

type SiteMap interface {
	Build(ctx *base.RequestContext, rooUrl string, depth int) interface{}
	Init()
}

type SiteMapImpl struct{}

func (s *SiteMapImpl) Build(ctx *base.RequestContext, rootUrl string, depth int) interface{} {
	root := graph.Build(ctx, rootUrl, depth)
	view := ViewWrapper(ctx, viewType, depth, root)
	return view
}

func (s *SiteMapImpl) Init(scrappingAlgo int, view int) {
	graph = new(GraphImpl)
	viewType = view
	graph.Init(scrappingAlgo)
}
