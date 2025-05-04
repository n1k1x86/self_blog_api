package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ArticleDB struct {
	ID          sql.NullInt64
	Title       sql.NullString
	ArticleText sql.NullString
	TagID       sql.NullInt64
	PublishedAt sql.NullTime
}

type ArticleREST struct {
	ID          int64      `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	ArticleText string     `json:"article_text,omitempty"`
	Tag         *TagREST   `json:"tag,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

type ArticleToSave struct {
	Title       string `json:"title,omitempty"`
	ArticleText string `json:"article_text,omitempty"`
	TagID       int64  `json:"tag_id,omitempty"`
}

func checkNullText(dbValue pgtype.Text) (string, error) {
	if dbValue.Valid {
		return dbValue.String, nil
	}
	return "", nil
}

func checkNullInt8(dbValue pgtype.Int8) (int64, error) {
	if dbValue.Valid {
		return dbValue.Int64, nil
	}
	return 0, nil
}

func checkNullTime(dbValue pgtype.Timestamp) (*time.Time, error) {
	if dbValue.Valid {
		return &dbValue.Time, nil
	}
	return nil, nil
}

func GetArticleByID(ctx context.Context, dbpool *pgxpool.Pool, id int64) (*ArticleREST, error) {
	query := `SELECT a.id, a.title, a.article_text, a.tag_id, a.published_at, t.name FROM articles as a INNER JOIN tags as t on t.id = a.tag_id where a.id = $1`
	res := dbpool.QueryRow(ctx, query, id)
	article := &ArticleREST{}
	tag := &TagREST{}

	var articleId pgtype.Int8
	var title pgtype.Text
	var articleText pgtype.Text
	var tagId pgtype.Int8
	var publishedAt pgtype.Timestamp
	var tagName pgtype.Text
	err := res.Scan(&articleId, &title, &articleText, &tagId, &publishedAt, &tagName)
	if err != nil {
		return nil, err
	}
	idValue, err := checkNullInt8(articleId)
	if err != nil {
		return nil, err
	}
	article.ID = idValue

	titleValue, err := checkNullText(title)
	if err != nil {
		return nil, err
	}
	article.Title = titleValue

	articleTextValue, err := checkNullText(articleText)
	if err != nil {
		return nil, err
	}
	article.ArticleText = articleTextValue

	publishedAtValue, err := checkNullTime(publishedAt)
	if err != nil {
		return nil, err
	}
	article.PublishedAt = publishedAtValue

	tagIDValue, err := checkNullInt8(tagId)
	if err != nil {
		return nil, err
	}
	tag.ID = tagIDValue

	tagNameValue, err := checkNullText(tagName)
	if err != nil {
		return nil, err
	}
	tag.Name = tagNameValue
	article.Tag = tag
	return article, nil
}

func GetArticles(ctx context.Context, dbpool *pgxpool.Pool) ([]ArticleREST, error) {
	query := `SELECT a.id::int8, a.title::text, a.article_text::text, a.tag_id::int8, a.published_at::timestamp, t.name::text FROM articles AS a LEFT JOIN tags AS t ON a.tag_id = t.id`
	res, err := dbpool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	articles := make([]ArticleREST, 0)

	for res.Next() {
		article := ArticleREST{}
		article.Tag = &TagREST{}
		var id pgtype.Int8
		var title pgtype.Text
		var articleText pgtype.Text
		var tagID pgtype.Int8
		var publisheAt pgtype.Timestamp
		var tagName pgtype.Text

		res.Scan(&id, &title, &articleText, &tagID, &publisheAt, &tagName)

		idValue, err := checkNullInt8(id)
		if err != nil {
			return nil, err
		}
		article.ID = idValue

		titleValue, err := checkNullText(title)
		if err != nil {
			return nil, err
		}
		article.Title = titleValue

		articleTextValue, err := checkNullText(articleText)
		if err != nil {
			return nil, err
		}
		article.ArticleText = articleTextValue

		tagIDValue, err := checkNullInt8(tagID)
		if err != nil {
			return nil, err
		}
		article.Tag.ID = tagIDValue

		publsAtValue, err := checkNullTime(publisheAt)
		if err != nil {
			return nil, err
		}
		article.PublishedAt = publsAtValue

		tagNameValue, err := checkNullText(tagName)
		if err != nil {
			return nil, err
		}
		article.Tag.Name = tagNameValue

		articles = append(articles, article)
	}
	if err = res.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

func AddArticle(ctx context.Context, dbpool *pgxpool.Pool, article *ArticleToSave) error {
	query := `INSERT INTO articles(title, article_text, tag_id) VALUES($1, $2, $3)`

	err := IsTagExist(ctx, dbpool, article.TagID)
	if err != nil {
		return err
	}
	_, err = dbpool.Exec(ctx, query, article.Title, article.ArticleText, article.TagID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(ctx context.Context, dbpool *pgxpool.Pool, id int64) error {
	query := `DELETE FROM articles WHERE id = $1`
	_, err := dbpool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func EditArticle(ctx context.Context, dbpool *pgxpool.Pool, article *ArticleToSave, id int64) error {
	query := `UPDATE articles SET title = $1, article_text = $2, tag_id = $3 WHERE id = $4`

	err := IsTagExist(ctx, dbpool, article.TagID)
	if err != nil {
		return err
	}
	_, err = dbpool.Exec(ctx, query, article.Title, article.ArticleText, article.TagID, id)
	if err != nil {
		return err
	}
	return nil
}
