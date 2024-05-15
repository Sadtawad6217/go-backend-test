package repository

import (
	"errors"
	"gobackend/pkg/core/model"
	"time"

	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("not found")

type Repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
func (r *Repository) GetPostAll(limit, offset int, searchTitle string) ([]model.Posts, error) {
	var articles []model.Posts
	query := `SELECT id, title, content, published, view_count, created_at, updated_at FROM posts WHERE published = true`
	if searchTitle != "" {
		query += ` AND title ILIKE '%' || $3 || '%'`
	}

	query += ` LIMIT $1 OFFSET $2`

	if searchTitle != "" {
		err := r.db.Select(&articles, query, limit, offset, searchTitle)
		if err != nil {
			return nil, err
		}
	} else {
		err := r.db.Select(&articles, query, limit, offset)
		if err != nil {
			return nil, err
		}
	}

	return articles, nil
}

func (r *Repository) GetTotalPostCount(searchTitle string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM posts WHERE published = true`
	if searchTitle != "" {
		query += ` AND title LIKE '%' || $1 || '%'`
		err := r.db.Get(&count, query, searchTitle)
		if err != nil {
			return 0, err
		}
	} else {
		err := r.db.Get(&count, query)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (r *Repository) GetPostID(id string) ([]model.Posts, error) {
	var articles []model.Posts
	query := `SELECT id, title, content, published, view_count, created_at FROM posts WHERE id = $1`
	err := r.db.Select(&articles, query, id)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *Repository) IncrementViewCount(id string) error {
	query := `UPDATE posts SET view_count = view_count + 1 WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreatePosts(post model.Posts) (model.Posts, error) {
	post.Published = false
	query := `INSERT INTO posts (title, content, Published) VALUES ($1, $2, $3) RETURNING id, title, content, published, created_at`
	row := r.db.QueryRow(query, post.Title, post.Content, post.Published)
	var createdArticle model.Posts
	err := row.Scan(&createdArticle.ID, &createdArticle.Title, &createdArticle.Content, &createdArticle.Published, &createdArticle.CreatedAt)
	if err != nil {
		return model.Posts{}, err
	}
	return createdArticle, nil
}

func (r *Repository) UpdatePost(id string, updateData model.Posts) (model.Posts, error) {
	query := `UPDATE posts 
			   SET title = $2, content = $3, published = $4, updated_at = $5 
			   WHERE id = $1 
			   RETURNING id, title, content, published, created_at`
	var updatedPost model.Posts
	err := r.db.Get(&updatedPost, query, id, updateData.Title, updateData.Content, updateData.Published, time.Now())
	if err != nil {
		return model.Posts{}, err
	}
	return updatedPost, nil
}

func (r *Repository) DeletePost(id string) error {
	query := "DELETE FROM posts WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
