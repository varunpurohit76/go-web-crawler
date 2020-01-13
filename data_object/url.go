package data_object

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/varunpurohit76/crawler/base"
)

type Url struct {
	Id   string
	Link string
}

type UrlDO interface {
	Get(ctx *base.RequestContext, tx *sqlx.Tx, id string) (*Url, error)
	Set(ctx *base.RequestContext, tx *sqlx.Tx, u *Url) (string, error)
	New(link string) *Url
}

type UrlImpl struct{}

func (u *UrlImpl) Get(ctx *base.RequestContext, tx *sqlx.Tx, parentId string) (*Url, error) {
	return GetUrlObj(ctx, tx, parentId)
}

func (u *UrlImpl) Set(ctx *base.RequestContext, tx *sqlx.Tx, urlObj *Url) (string, error) {
	return CreateUrlObj(ctx, tx, urlObj)
}

func (u *UrlImpl) New(link string) *Url {
	return &Url{
		Id:   uuid.New().String(),
		Link: link,
	}
}
