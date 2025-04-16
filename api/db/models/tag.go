package models

import "database/sql"

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
