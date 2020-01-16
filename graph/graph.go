package graph

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/varunpurohit76/crawler/base"
	. "github.com/varunpurohit76/crawler/data_object"
	. "github.com/varunpurohit76/crawler/scrapper"
)

type Graph interface {
	Build(ctx *base.RequestContext, rootUrl string, depth int) *Url
	Init(algorithm int)
}

var urlDO UrlDO
var relationDO RelationDO
var scrapper Scrapper

type GraphImpl struct{}

func (g *GraphImpl) Build(ctx *base.RequestContext, rootUrl string, depth int) *Url {
	defer base.LogLatency("sitemap.graph.build.latency", nil, time.Now())
	rootUrlObj := urlDO.New(rootUrl)
	_, err := urlDO.Set(ctx, nil, rootUrlObj)
	if err != nil {
		return nil
	}
	graphBuildRecursion(ctx, rootUrlObj, depth)
	return rootUrlObj
}

func graphBuildRecursion(ctx *base.RequestContext, rootUrlObj *Url, depth int) {
	if depth == 0 {
		return
	}
	ctx.Logger().WithFields(logrus.Fields{"rootUrlObjLink": rootUrlObj.Link, "rootUrlObjId": rootUrlObj.Id, "depth": depth}).Debug("graph build")

	// 1. call scrapper for children url
	childrenUlr, err := scrapper.Scrap(ctx, rootUrlObj.Link)
	if err != nil {
		return
	}

	// 2. create children urlDO
	var children []*Url
	for _, c := range childrenUlr {
		child := urlDO.New(c)
		children = append(children, child)
	}

	// 3. persist children urlDO
	var childrenId []string
	for _, c := range children {
		childId, err := urlDO.Set(ctx, nil, c)
		if err != nil {
			return
		}
		childrenId = append(childrenId, childId)
	}

	// 4. create and persist root <-> children relationDO
	for _, cId := range childrenId {
		err := relationDO.Set(ctx, nil, relationDO.New(rootUrlObj.Id, cId))
		if err != nil {
			return
		}
	}

	// 5. recursively call Build for all children
	var wg sync.WaitGroup
	wg.Add(len(children))
	for _, child := range children {
		go func(child *Url) {
			defer wg.Done()
			graphBuildRecursion(ctx, child, depth-1)
		}(child)
	}
	wg.Wait()

	return
}

func (g *GraphImpl) Init(algorithm int) {
	urlDO = new(UrlImpl)
	relationDO = new(RelationImpl)
	scrapper = ScrapperFactory(algorithm)
}
