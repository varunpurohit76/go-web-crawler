package scrapper

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/varunpurohit76/crawler/base"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Scrapper interface {
	Scrap(ctx *base.RequestContext, rootUrl string) (childUrl []string, err error)
}

type urlExtract struct{}
type mockUrlExtract struct {}
type void struct{}

var member void

const (
	PageUrlExtract = iota
	MockPageUrlExtract
)

func ScrapperFactory(algo int) Scrapper {
	switch algo {
	case PageUrlExtract:
		return new(urlExtract)
	case MockPageUrlExtract:
		return new(mockUrlExtract)
	default:
		return nil
	}
}

func (s *urlExtract) Scrap(ctx *base.RequestContext, rootUrl string) ([]string, error) {
	defer base.LogLatency("scrap.latency", logrus.Fields{"rootUrl": rootUrl}, time.Now())
	var links []string
	var linksSet = make(map[string]void)
	baseUrl, err := url.Parse(rootUrl)
	if err != nil {
		return links, err
	}
	resp, _ := http.Get(rootUrl)
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					link := formatURL(baseUrl, attr.Val)
					linksSet[link] = member
				}
			}
		}
	}
	delete(linksSet, rootUrl)
	for k, _ := range linksSet {
		if strings.Contains(k, rootUrl) {
			links = append(links, k)
		}
	}
	ctx.Logger().WithFields(logrus.Fields{"rootUrl": rootUrl}).Debug(fmt.Sprintf("found %v links", len(links)))
	return links, nil
}

func formatURL(baseUrl *url.URL, link string) string {
	linkURL, err := url.Parse(link)
	if err != nil {
		return ""
	}
	uriFormatted := baseUrl.ResolveReference(linkURL)
	linkStr := uriFormatted.String()
	if strings.HasSuffix(link, "/") {
		linkStr = linkStr[:len(linkStr)-1]
	}
	return linkStr
}

func (m *mockUrlExtract) Scrap(ctx *base.RequestContext, rootUrl string) ([]string, error) {
	defer base.LogLatency("scrap.latency", logrus.Fields{"rootUrl": rootUrl}, time.Now())
	if rootUrl == "https://www.monzo.com" {
		return []string{"https://www.monzo.com/a", "https://www.monzo.com/b"}, nil
	}
	if rootUrl == "https://www.monzo.com/a" {
		return []string{"https://www.monzo.com/a1", "https://www.monzo.com/a2"}, nil
	}
	if rootUrl == "https://www.monzo.com/a1" {
		return []string{"https://www.monzo.com/a11", "https://www.monzo.com/a12"}, nil
	}
	if rootUrl == "https://www.monzo.com/a2" {
		return []string{"https://www.monzo.com/a21", "https://www.monzo.com/a22"}, nil
	}
	if rootUrl == "https://www.monzo.com/b" {
		return []string{"https://www.monzo.com/b1", "https://www.monzo.com/b2"}, nil
	}
	if rootUrl == "https://www.monzo.com/b1" {
		return []string{"https://www.monzo.com/b11", "https://www.monzo.com/b12"}, nil
	}
	if rootUrl == "https://www.monzo.com/b2" {
		return []string{"https://www.monzo.com/b21", "https://www.monzo.com/b22"}, nil
	}
	return []string{}, nil
}