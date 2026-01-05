package endpoint

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/laiambryant/tcgdex/client"
	"github.com/laiambryant/tcgdex/query"
)

type Endpoint[T any, L any] struct {
	Client *client.Client
	Path   string
}

func New[T any, L any](c *client.Client, path string) *Endpoint[T, L] {
	return &Endpoint[T, L]{
		Client: c,
		Path:   path,
	}
}

func (e *Endpoint[T, L]) Get(ctx context.Context, id string) (T, error) {
	var item T
	path := fmt.Sprintf("/%s/%s", e.Path, id)
	data, err := e.Client.Get(ctx, path)
	if err != nil {
		return item, err
	}
	if err := json.Unmarshal(data, &item); err != nil {
		return item, &DecodeError{Resource: path, Err: err}
	}
	return item, nil
}

func (e *Endpoint[T, L]) List(ctx context.Context, q *query.Query) ([]L, error) {
	var items []L
	qs := ""
	if q != nil {
		qs = q.Build()
	}
	path := fmt.Sprintf("/%s%s", e.Path, qs)
	data, err := e.Client.Get(ctx, path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, &DecodeError{Resource: path, Err: err}
	}
	return items, nil
}
