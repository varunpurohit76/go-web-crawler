package data_object

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type relationDBO struct {
	Parent sql.NullString
	Child  sql.NullString
}

func createRelationDb(tx *sqlx.Tx, r *Relation) (error) {
	if tx == nil {
		return errors.New("no tx")
	}
	parentId := sql.NullString{
		String: r.ParentId,
		Valid:  r.ParentId != "",
	}
	childId := sql.NullString{
		String: r.ChildId,
		Valid:  r.ChildId != "",
	}
	_, err := tx.Exec("INSERT INTO relation (parent, child) VALUES(?, ?)", parentId, childId)
	if err != nil {
		return err
	}
	return nil
}

func getRelationDb(tx *sqlx.Tx, parentId string) ([]*Relation, error) {
	if tx == nil {
		return nil, errors.New("no tx")
	}
	relationDBOs := []relationDBO{}
	err := tx.Select(&relationDBOs, "SELECT * FROM relation WHERE parent=?", parentId)
	if err != nil {
		return nil, err
	}
	relations := make([]*Relation, 0)
	for _, r := range relationDBOs {
		newRelation := &Relation{
			ParentId: r.Parent.String,
			ChildId:  r.Child.String,
		}
		relations = append(relations, newRelation)
	}
	return relations, nil
}
