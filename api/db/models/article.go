package models

import (
	"database/sql"
	"time"
)

type ArticleDB struct {
	ID          sql.NullInt64
	Title       sql.NullString
	ArticleText sql.NullString
	TagID       sql.NullInt64
	PublishedAt sql.NullTime
}

type ArticleREST struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	ArticleText string    `json:"article_text"`
	TagID       int64     `json:"tag_id"`
	PublishedAt time.Time `json:"published_at"`
}

func ArticleDBToREST(article *ArticleDB) (*ArticleREST, error) {
	var id int64
	var title string
	var articleText string
	var tagID int64
	var publisheAt time.Time

	if article.ID.Valid {
		idDB, err := article.ID.Value()
		if err != nil {
			return nil, err
		}
		id = idDB.(int64)
	} else {
		id = 0
	}

	if article.Title.Valid {
		titleDB, err := article.Title.Value()
		if err != nil {
			return nil, err
		}
		title = titleDB.(string)
	} else {
		title = ""
	}

	if article.ArticleText.Valid {
		articleTextDB, err := article.Title.Value()
		if err != nil {
			return nil, err
		}
		articleText = articleTextDB.(string)
	} else {
		articleText = ""
	}

	if article.TagID.Valid {
		tagIDDB, err := article.TagID.Value()
		if err != nil {
			return nil, err
		}
		tagID = tagIDDB.(int64)
	} else {
		tagID = 0
	}

	if article.PublishedAt.Valid {
		publisheAtDB, err := article.PublishedAt.Value()
		if err != nil {
			return nil, err
		}
		publisheAt = publisheAtDB.(time.Time)
	} else {
		publisheAt = time.Now()
	}

	return &ArticleREST{
		ID:          id,
		Title:       title,
		ArticleText: articleText,
		TagID:       tagID,
		PublishedAt: publisheAt,
	}, nil
}
