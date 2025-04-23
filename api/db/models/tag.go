package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
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

func AddTag(ctx context.Context, dbpool *pgxpool.Pool, tag *TagREST) error {
	query := fmt.Sprintf(`INSERT INTO tags(name) VALUES('%s') RETURNING id`, tag.Name)
	var tagID int64
	err := dbpool.QueryRow(ctx, query).Scan(&tagID)
	if err != nil {
		return err
	}
	tag.ID = tagID
	return nil
}

func GetTags(ctx context.Context, dbpool *pgxpool.Pool) ([]TagREST, error) {
	query := `SELECT * FROM tags`
	resRows, err := dbpool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var tags []TagREST

	for resRows.Next() {
		var id int64
		var name string

		resRows.Scan(&id, &name)
		log.Println(id, name)
		tags = append(tags, TagREST{ID: id, Name: name})
	}
	log.Println(tags)

	return tags, nil
}

func GetTagByID(ctx context.Context, dbpool *pgxpool.Pool, id int64) (*TagREST, error) {
	query := fmt.Sprintf(`SELECT id, name FROM tags WHERE id = %d`, id)
	rows, err := dbpool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var name sql.NullString

	for rows.Next() {
		rows.Scan(&id, &name)
	}
	tag := &TagREST{}

	if name.Valid {
		value, err := name.Value()
		if err != nil {
			return nil, err
		}
		tag.ID = id
		tag.Name = value.(string)
	} else {
		return nil, ErrNotFound
	}
	return tag, nil
}

func DeleteTag(ctx context.Context, dbpool *pgxpool.Pool, id int64) error {
	query := fmt.Sprintf(`DELETE FROM tags WHERE id = %d`, id)
	res, err := dbpool.Exec(ctx, query)
	if err != nil {
		return err
	}
	cnt := res.RowsAffected()
	if cnt == 0 {
		return ErrNotFound
	}
	return nil
}

func EditTag(ctx context.Context, dbpool *pgxpool.Pool, tag *TagREST) error {
	query := fmt.Sprintf(`UPDATE tags SET name = '%s' WHERE id = %d`, tag.Name, tag.ID)
	res, err := dbpool.Exec(ctx, query)
	if err != nil {
		return err
	}
	cnt := res.RowsAffected()
	if cnt == 0 {
		return ErrNotFound
	}
	return nil
}
