package graph

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/varunpurohit76/crawler/base"
	. "github.com/varunpurohit76/crawler/data_object"
)

type SiteMapUrlNodeView struct {
	Link     string                `json:"link"`
	Children []*SiteMapUrlNodeView `json:"children"`
}

const (
	JsonView = iota
)

func ViewWrapper(ctx *base.RequestContext, kind int, depth int, root *Url) interface{} {
	switch kind {
	case JsonView:
		defer base.LogLatency("sitemap.graph.view.latency", nil, time.Now())
		return jsonView(ctx, depth, root)
	default:
		return nil
	}
}

func jsonView(ctx *base.RequestContext, depth int, root *Url) *SiteMapUrlNodeView {
	ctx.Logger().WithFields(logrus.Fields{"rootUrl": root.Link, "rootId": root.Id, "depth": depth}).Debug("sitemap view build")
	var ch []*SiteMapUrlNodeView
	if depth > 0 {
		rel, err := relationDO.Get(ctx, nil, root.Id)
		if err != nil {
			return nil
		}
		ctx.Logger().Debug(fmt.Sprintf("found %v relations for %v", len(rel), root.Link))
		var wg sync.WaitGroup
		wg.Add(len(rel))
		children := make(chan *SiteMapUrlNodeView, len(rel))
		for _, r := range rel {
			go func(r *Relation) {
				defer wg.Done()
				child, err := urlDO.Get(ctx, nil, r.ChildId)
				if err != nil && child != nil {
					children <- nil
				}
				children <- jsonView(ctx, depth-1, child)
			}(r)
		}
		wg.Wait()
		close(children)
		for c := range children {
			ch = append(ch, &SiteMapUrlNodeView{
				Link:     c.Link,
				Children: c.Children,
			})
		}
	}
	return &SiteMapUrlNodeView{
		Link:     root.Link,
		Children: ch,
	}
}
