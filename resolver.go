//go:generate gorunpkg github.com/99designs/gqlgen

package writing

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type Resolver struct{}

func New() Config {
	c := Config{
		Resolvers: &Resolver{},
	}
	return c
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreatePost(ctx context.Context, input NewPost) (Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) EditPost(ctx context.Context, id string, input NewPost) (Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLink(ctx context.Context, input NewLink) (Link, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) AllPosts(ctx context.Context) ([]*Post, error) {
	rows, err := db.Query("SELECT id, title, content, date, created_at, modified_at, tags, draft FROM posts WHERE draft = false ORDER BY date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]*Post, 0)
	for rows.Next() {
		post := new(Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Datetime, &post.Created, &post.Modified, pq.Array(&post.Tags), &post.Draft)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*Post, error) {
	rows, err := db.Query("SELECT id, title, content, date, created_at, modified_at, tags, draft FROM posts WHERE draft = false ORDER BY date DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]*Post, 0)
	for rows.Next() {
		post := new(Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Datetime, &post.Created, &post.Modified, pq.Array(&post.Tags), &post.Draft)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*Post, error) {
	var post Post
	row := db.QueryRow("SELECT id, title, content, date, created_at, modified_at, tags, draft FROM posts WHERE id = $1", id)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Datetime, &post.Created, &post.Modified, pq.Array(&post.Tags), &post.Draft)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("No post with id %s", id)
	case err != nil:
		return nil, fmt.Errorf("Error running get query: %+v", err)
	default:
		return &post, nil
	}
}

func (r *queryResolver) AllLinks(ctx context.Context) ([]*Link, error) {
	panic("not implemented")
}
func (r *queryResolver) Links(ctx context.Context, limit *int, offset *int) ([]*Link, error) {
	panic("not implemented")
}
func (r *queryResolver) Link(ctx context.Context, id string) (*Link, error) {
	panic("not implemented")
}