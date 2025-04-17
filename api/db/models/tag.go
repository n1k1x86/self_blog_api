package models

import (
	"api/config"
	"api/db/dbconns"
	"database/sql"
	"encoding/json"
	"fmt"
)

type TagDB struct {
	ID   sql.NullInt64
	Name sql.NullString
}

type TagREST struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func TagDBToREST(tag *TagDB) (*TagREST, error) {
	var id int64
	var name string

	if tag.ID.Valid {
		idDB, err := tag.ID.Value()
		if err != nil {
			return nil, err
		}
		id = idDB.(int64)
	} else {
		id = 0
	}
	if tag.Name.Valid {
		nameDB, err := tag.Name.Value()
		if err != nil {
			return nil, err
		}
		name = nameDB.(string)
	} else {
		name = ""
	}
	return &TagREST{
		ID:   id,
		Name: name,
	}, nil
}

func AddTag(cfg config.BlogDBConfig, body []byte) (int64, error) {
	db, err := dbconns.ConnectToBlogDBConfig(cfg)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	tagrest := &TagREST{}
	err = json.Unmarshal(body, tagrest)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf(`INSERT INTO tags(name) VALUES('%s') RETURNING id`, tagrest.Name)
	var tagID int64
	err = db.QueryRow(query).Scan(&tagID)
	if err != nil {
		return 0, err
	}
	return tagID, nil
}

func GetTags(cfg config.BlogDBConfig) ([]byte, error) {
	db, err := dbconns.ConnectToBlogDBConfig(cfg)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `SELECT * FROM tags`
	resRows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var tags []TagREST

	for resRows.Next() {
		var id int64
		var name string

		resRows.Scan(&id, &name)
		tags = append(tags, TagREST{ID: id, Name: name})
	}

	result, err := json.Marshal(tags)
	if err != nil {
		return nil, err
	}
	return result, nil
}
