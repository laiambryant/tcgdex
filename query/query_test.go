package query

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestBuildEmpty(t *testing.T) {
	q := New()
	if got := q.Build(); got != "" {
		t.Fatalf("expected empty string for empty query, got %q", got)
	}
}

func TestChainingReturnsSamePointer(t *testing.T) {
	q := New()
	q2 := q.Contains("name", "bob")
	if q != q2 {
		t.Fatalf("expected methods to return same pointer receiver")
	}
}

func TestAllMethodsBuildsExpectedQuery(t *testing.T) {
	q := New()
	q.Contains("name", "bob")
	q.Equal("status", "active")
	q.NotEqual("role", "admin")
	q.GTE("score", 10)
	q.LTE("rank", 5)
	q.GT("visits", 100)
	q.LT("errors", 2)
	q.IsNull("deleted_at")
	q.NotNull("created_at")
	q.NotContains("desc", "spoiler")
	q.Sort("title", "desc")
	q.Paginate(3, 25)
	pairs := [][2]string{
		{"name", "bob"},
		{"status", "eq:active"},
		{"role", "neq:admin"},
		{"score", "gte:10"},
		{"rank", "lte:5"},
		{"visits", "gt:100"},
		{"errors", "lt:2"},
		{"deleted_at", "null:"},
		{"created_at", "notnull:"},
		{"desc", "not:spoiler"},
		{"sort:field", "title"},
		{"sort:order", "desc"},
		{"pagination:page", "3"},
		{"pagination:itemsPerPage", "25"},
	}
	var parts []string
	for _, p := range pairs {
		parts = append(parts, fmt.Sprintf("%s=%s", url.QueryEscape(p[0]), url.QueryEscape(p[1])))
	}
	expected := "?" + strings.Join(parts, "&")
	if got := q.Build(); got != expected {
		t.Fatalf("unexpected query string:\n got: %s\nwant: %s", got, expected)
	}
}

func TestQueryEscaping(t *testing.T) {
	q := New()
	q.Contains("a b", "c&d=!")
	expected := "?" + strings.Join([]string{fmt.Sprintf("%s=%s", url.QueryEscape("a b"), url.QueryEscape("c&d=!"))}, "&")
	if got := q.Build(); got != expected {
		t.Fatalf("escaping mismatch: got %q want %q", got, expected)
	}
}
