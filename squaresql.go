package squaresql

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
)

// Preparer is an interface used by Prepare.
type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

// PreparerContext is an interface used by PrepareContext.
type PreparerContext interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// Queryer is an interface used by Query.
type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// QueryerContext is an interface used by QueryContext.
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// QueryRower is an interface used by QueryRow.
type QueryRower interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

// QueryRowerContext is an interface used by QueryRowContext.
type QueryRowerContext interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// Execer is an interface used by Exec.
type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// ExecerContext is an interface used by ExecContext.
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type SquareSql struct {
	queries map[string]string
}

func (s *SquareSql) lookupQuery(name string) (query string, err error) {
	query, ok := s.queries[name]
	if !ok {
		err = fmt.Errorf("dotsql: '%s' could not be found", name)
	}

	return
}

func (s *SquareSql) Prepare(db Preparer, name string) (*sql.Stmt, error) {
	query, err := s.lookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.Prepare(query)
}

func (s *SquareSql) PrepareContext(ctx context.Context, db PreparerContext, name string) (*sql.Stmt, error) {
	query, err := s.lookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.PrepareContext(ctx, query)
}

func (s *SquareSql) Query(db Queryer, name string, args ...interface{}) (*sql.Rows, error) {
	query, err := s.lookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.Query(query, args...)
}

func (s *SquareSql) QueryContext(ctx context.Context, db QueryerContext, name string, args ...interface{}) (*sql.Rows, error) {
	query, err := s.lookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.QueryContext(ctx, query, args...)
}

func (s *SquareSql) QueryRow(db QueryRower, name string, args ...interface{}) (*sql.Row, error) {
	query, err := s.lookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.QueryRow(query, args...), nil
}

func (s *SquareSql) QueryRowContext(ctx context.Context, db QueryRowerContext, name string, args ...interface{}) (*sql.Row, error) {
	query, err := s.lookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.QueryRowContext(ctx, query, args...), nil
}

func (s *SquareSql) Exec(db Execer, name string, args ...interface{}) (sql.Result, error) {
	query, err := s.lookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.Exec(query, args...)
}

func (s *SquareSql) ExecContext(ctx context.Context, db ExecerContext, name string, args ...interface{}) (sql.Result, error) {
	query, err := s.lookupQuery(name)
	if err != nil {
		return nil, err
	}

	return db.ExecContext(ctx, query, args...)
}

func (s *SquareSql) Raw(name string) (string, error) {
	return s.lookupQuery(name)
}

func (s *SquareSql) QueryMap() map[string]string {
	return s.queries
}

func Load(r io.Reader) (*SquareSql, error) {
	scanner := &Scanner{}
	queries := scanner.Run(bufio.NewScanner(r))
	squaresql := &SquareSql{queries: queries}
	return squaresql, nil
}

func LoadFromString(sql string) (*SquareSql, error) {
	buf := bytes.NewBufferString(sql)
	return Load(buf)
}

func LoadFromFile(sqlFile string) (*SquareSql, error) {
	f, err := os.Open(sqlFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Load(f)
}

func Merge(dots ...*SquareSql) *SquareSql {
	queries := make(map[string]string)

	for _, dot := range dots {
		for k, v := range dot.QueryMap() {
			queries[k] = v
		}
	}

	return &SquareSql{
		queries: queries,
	}
}
