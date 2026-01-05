package query

import (
"fmt"
"net/url"
"strings"
)

type Query struct {
	params []param
}

type param struct {
	key   string
	value string
}

func New() *Query {
	return &Query{}
}

func (q *Query) add(key, value string) {
	q.params = append(q.params, param{key, value})
}

func (q *Query) Contains(key, value string) *Query {
	q.add(key, value)
	return q
}

func (q *Query) Equal(key, value string) *Query {
	q.add(key, "eq:"+value)
	return q
}

func (q *Query) NotEqual(key, value string) *Query {
	q.add(key, "neq:"+value)
	return q
}

func (q *Query) GTE(key string, value int) *Query {
	q.add(key, fmt.Sprintf("gte:%d", value))
	return q
}

func (q *Query) LTE(key string, value int) *Query {
	q.add(key, fmt.Sprintf("lte:%d", value))
	return q
}

func (q *Query) GT(key string, value int) *Query {
	q.add(key, fmt.Sprintf("gt:%d", value))
	return q
}

func (q *Query) LT(key string, value int) *Query {
	q.add(key, fmt.Sprintf("lt:%d", value))
	return q
}

func (q *Query) IsNull(key string) *Query {
	q.add(key, "null:")
	return q
}

func (q *Query) NotNull(key string) *Query {
	q.add(key, "notnull:")
	return q
}

func (q *Query) NotContains(key, value string) *Query {
	q.add(key, "not:"+value)
	return q
}

func (q *Query) Sort(field, order string) *Query {
	q.add("sort:field", field)
	q.add("sort:order", order)
	return q
}

func (q *Query) Paginate(page, itemsPerPage int) *Query {
	q.add("pagination:page", fmt.Sprintf("%d", page))
	q.add("pagination:itemsPerPage", fmt.Sprintf("%d", itemsPerPage))
	return q
}

func (q *Query) Build() string {
	if len(q.params) == 0 {
		return ""
	}
	var parts []string
	for _, p := range q.params {
		k := url.QueryEscape(p.key)
		v := url.QueryEscape(p.value)
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}
	return "?" + strings.Join(parts, "&")
}
