package squaresql

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetTag(t *testing.T) {
	var tests = []struct {
		line string
		want string
	}{
		{"SELECT all", ""},
		{"-- no name", ""},
		{"-- name:  ", ""},
		{"-- name: find-products-by-name", "find-products-by-name"},
		{"  --  name:  save-product ", "save-product"},
	}

	for _, c := range tests {
		got := getTag(c.line)
		assert.Equal(t, got, c.want)
	}
}

func TestRun(t *testing.T) {
	sqlFile := `
	-- name: all-products
	-- Finds all products
	SELECT * from products
	-- name: empty-query-should-not-be-stored
	-- name: save-products
	INSERT INTO products (?, ?, ?)
	`

	exp := map[string]string{
		"all-products":  "-- Finds all products\nSELECT * from products",
		"save-products": "INSERT INTO products (?, ?, ?)",
	}

	scanner := &Scanner{}
	queries := scanner.Run(bufio.NewScanner(strings.NewReader(sqlFile)))

	assert.Equal(t, queries, exp)
}
