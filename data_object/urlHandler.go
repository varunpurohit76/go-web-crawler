package data_object

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/varunpurohit76/crawler/base"
)

func GetUrlObj(ctx *base.RequestContext, tx *sqlx.Tx, id string) (*Url, error) {
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

	url, err := getUrlDb(tx, id)
	if err != nil {
		ctx.Logger().WithField("id", id).Error("url db get fail")
		return nil, err
	}

	if txCreated {
		err := tx.Commit()
		if err != nil {
			return nil, err
		}
	}
	ctx.Logger().WithFields(logrus.Fields{"id": url.Id, "link": url.Link}).Info("url db get success")
	return url, nil
}

func CreateUrlObj(ctx *base.RequestContext, tx *sqlx.Tx, urlObj *Url) (string, error) {
	var err error
	txCreated := false
	if tx == nil {
		tx, err = base.NewDbTransaction()
		if err != nil {
			return "", err
		}
		defer tx.Rollback()
		txCreated = true
	}

	id, err := createUrlDb(tx, urlObj)
	if err != nil {
		ctx.Logger().WithFields(logrus.Fields{"id": urlObj.Id, "link": urlObj.Link}).Error("url db create fail")
		return "", err
	}

	if txCreated {
		err := tx.Commit()
		if err != nil {
			return "", err
		}
	}
	ctx.Logger().WithFields(logrus.Fields{"id": urlObj.Id, "link": urlObj.Link}).Info("url db create success")
	return id, nil
}
