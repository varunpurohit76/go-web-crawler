package data_object

import (
	"github.com/jmoiron/sqlx"
	"github.com/varunpurohit76/crawler/base"
)

type Relation struct {
	ParentId string
	ChildId  string
}

type RelationDO interface {
	Get(ctx *base.RequestContext, tx *sqlx.Tx, parentId string) ([]*Relation, error)
	Set(ctx *base.RequestContext, tx *sqlx.Tx, r *Relation) error
	New(parentId string, childId string) *Relation
}

type RelationImpl struct{}

func (r *RelationImpl) Get(ctx *base.RequestContext, tx *sqlx.Tx, parentId string) ([]*Relation, error) {
	return GetRelationObj(ctx, tx, parentId)
}

func (r *RelationImpl) Set(ctx *base.RequestContext, tx *sqlx.Tx, rel *Relation) error {
	return CreateRelationObj(ctx, tx, rel)
}

func (r *RelationImpl) New(parentId string, childId string) *Relation {
	return &Relation{
		ParentId: parentId,
		ChildId:  childId,
	}
}
