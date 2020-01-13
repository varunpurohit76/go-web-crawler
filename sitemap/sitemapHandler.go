package sitemap

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/varunpurohit76/crawler/base"
)

type SiteMapRequest struct {
	Url   string `json:"url"`
	Depth int    `json:"depth"`
}

func SitemapHandler(w http.ResponseWriter, r *http.Request) {
	req := &SiteMapRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rc := base.NewRequestContext()
	rc.Logger().WithFields(logrus.Fields{"url": req.Url, "depth": req.Depth}).Info("siteMapBuild init")
	res := Service.Build(rc, req.Url, req.Depth)
	resJson, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resJson)
}
