package data_object

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/varunpurohit76/crawler/base"
)

func GetRelationObj(ctx *base.RequestContext, tx *sqlx.Tx, parentId string) ([]*Relation, error) {
	var err error
	txCreated := false
	if tx == nil {
		tx, err = base.NewDbTransaction()
		if err != nil {
			return nil, err
		}
		defer tx.Rollback()
		txCreated = true
	}

	relations, err := getRelationDb(tx, parentId)
	if err != nil {
		ctx.Logger().WithField("parent_id", parentId).Error("relation db get fail")
		return nil, err
	}

	if txCreated {
		err := tx.Commit()
		if err != nil {
			return nil, err
		}
	}
	ctx.Logger().WithFields(logrus.Fields{"parent_id": parentId, "relations": len(relations)}).Debug("relation db get success")
	return relations, nil
}

func CreateRelationObj(ctx *base.RequestContext, tx *sqlx.Tx, relationObj *Relation) error {
	var err error
	txCreated := false
	if tx == nil {
		tx, err = base.NewDbTransaction()
		if err != nil {
			return err
		}
		defer tx.Rollback()
		txCreated = true
	}

	err = createRelationDb(tx, relationObj)
	if err != nil {
		ctx.Logger().WithFields(logrus.Fields{"parent_id": relationObj.ParentId, "child_id": relationObj.ChildId}).Error("relation db create fail")
		return err
	}

	if txCreated {
		err := tx.Commit()
		if err != nil {
			return err
		}
	}
	ctx.Logger().WithFields(logrus.Fields{"parent_id": relationObj.ParentId, "child_id": relationObj.ChildId}).Debug("relation db create success")
	return nil
}
