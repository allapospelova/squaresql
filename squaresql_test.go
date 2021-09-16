package squaresql

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPrepare(t *testing.T) {
	square := SquareSql{
		queries: map[string]string{"select": "SELECT * from products"},
	}

	preparerStub := func(err error) *PreparerMock {
		return &PreparerMock{
			PrepareFunc: func(_ string) (*sql.Stmt, error) {
				if err != nil {
					return nil, err
				}
				return &sql.Stmt{}, nil
			},
		}
	}

	tests := []struct {
		name          string
		square        SquareSql
		preparerStub  func(err error) *PreparerMock
		prCallsNumber int
		psArg         error
		prepareArg    string
		wantErr       bool
	}{
		{
			name:          "not found",
			square:        square,
			preparerStub:  preparerStub,
			psArg:         nil,
			wantErr:       true,
			prepareArg:    "insert",
			prCallsNumber: 0,
		},
		{
			name:          "error returned by db",
			square:        square,
			preparerStub:  preparerStub,
			psArg:         errors.New("critical error"),
			wantErr:       true,
			prepareArg:    "select",
			prCallsNumber: 1,
		},
		{
			name:          "successful query preparation",
			square:        square,
			preparerStub:  preparerStub,
			psArg:         nil,
			wantErr:       false,
			prepareArg:    "select",
			prCallsNumber: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.preparerStub(tt.psArg)
			stmt, err := tt.square.Prepare(p, tt.prepareArg)
			prCallsNumber := p.CallNumber()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, stmt)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stmt)
			}

			assert.Equal(t, prCallsNumber, tt.prCallsNumber)
		})
	}
}

func TestPrepareContext(t *testing.T) {
	square := SquareSql{
		queries: map[string]string{"select": "SELECT * from products"},
	}

	preparerContextStub := func(err error) *PreparerContextMock {
		return &PreparerContextMock{
			PrepareContextFunc: func(_ context.Context, _ string) (*sql.Stmt, error) {
				if err != nil {
					return nil, err
				}
				return &sql.Stmt{}, nil
			},
		}
	}

	ctx := context.Background()

	tests := []struct {
		name          string
		square        SquareSql
		preparerStub  func(err error) *PreparerContextMock
		prCallsNumber int
		psArg         error
		prepareArg    string
		wantErr       bool
	}{
		{
			name:          "not found",
			square:        square,
			preparerStub:  preparerContextStub,
			psArg:         nil,
			wantErr:       true,
			prepareArg:    "insert",
			prCallsNumber: 0,
		},
		{
			name:          "error returned by db",
			square:        square,
			preparerStub:  preparerContextStub,
			psArg:         errors.New("critical error"),
			wantErr:       true,
			prepareArg:    "select",
			prCallsNumber: 1,
		},
		{
			name:          "successful query preparation",
			square:        square,
			preparerStub:  preparerContextStub,
			psArg:         nil,
			wantErr:       false,
			prepareArg:    "select",
			prCallsNumber: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.preparerStub(tt.psArg)
			stmt, err := tt.square.PrepareContext(ctx, p, tt.prepareArg)
			prCallsNumber := p.CallNumber()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, stmt)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stmt)
			}

			assert.Equal(t, prCallsNumber, tt.prCallsNumber)
		})
	}
}

func TestQuery(t *testing.T) {
	square := SquareSql{
		queries: map[string]string{"select": "SELECT * from products WHERE id = ?"},
	}

	queryerStub := func(err error) *QueryerMock {
		return &QueryerMock{
			QueryFunc: func(_ string, _ ...interface{}) (*sql.Rows, error) {
				if err != nil {
					return nil, err
				}
				return &sql.Rows{}, nil
			},
		}
	}

	tests := []struct {
		name          string
		square        SquareSql
		queryerStub   func(err error) *QueryerMock
		prCallsNumber int
		psArg         error
		prepareArg    string
		wantErr       bool
	}{
		{
			name:          "error returned by db",
			square:        square,
			queryerStub:   queryerStub,
			psArg:         errors.New("critical error"),
			wantErr:       true,
			prepareArg:    "select",
			prCallsNumber: 1,
		},
		{
			name:          "not found",
			square:        square,
			queryerStub:   queryerStub,
			psArg:         nil,
			wantErr:       true,
			prepareArg:    "insert",
			prCallsNumber: 0,
		},
		{
			name:          "successful query preparation",
			square:        square,
			queryerStub:   queryerStub,
			psArg:         nil,
			wantErr:       false,
			prepareArg:    "select",
			prCallsNumber: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.queryerStub(tt.psArg)
			stmt, err := tt.square.Query(p, tt.prepareArg, "1")
			prCallsNumber := p.CallNumber()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, stmt)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stmt)
			}

			assert.Equal(t, prCallsNumber, tt.prCallsNumber)
		})
	}
}

func TestQueryContext(t *testing.T) {
	square := SquareSql{
		queries: map[string]string{"select": "SELECT * from products WHERE id = ?"},
	}

	queryerContextStub := func(err error) *QueryerContextMock {
		return &QueryerContextMock{
			QueryContextFunc: func(_ context.Context, _ string, _ ...interface{}) (*sql.Rows, error) {
				if err != nil {
					return nil, err
				}
				return &sql.Rows{}, nil
			},
		}
	}

	ctx := context.Background()

	tests := []struct {
		name               string
		square             SquareSql
		queryerContextStub func(err error) *QueryerContextMock
		prCallsNumber      int
		psArg              error
		prepareArg         string
		wantErr            bool
	}{
		{
			name:               "error returned by db",
			square:             square,
			queryerContextStub: queryerContextStub,
			psArg:              errors.New("critical error"),
			wantErr:            true,
			prepareArg:         "select",
			prCallsNumber:      1,
		},
		{
			name:               "not found",
			square:             square,
			queryerContextStub: queryerContextStub,
			psArg:              nil,
			wantErr:            true,
			prepareArg:         "insert",
			prCallsNumber:      0,
		},
		{
			name:               "successful query preparation",
			square:             square,
			queryerContextStub: queryerContextStub,
			psArg:              nil,
			wantErr:            false,
			prepareArg:         "select",
			prCallsNumber:      1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.queryerContextStub(tt.psArg)
			stmt, err := tt.square.QueryContext(ctx, p, tt.prepareArg, "1")
			prCallsNumber := p.CallNumber()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, stmt)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stmt)
			}

			assert.Equal(t, prCallsNumber, tt.prCallsNumber)
		})
	}
}

func TestQueryRow(t *testing.T) {
	square := SquareSql{
		queries: map[string]string{"select": "SELECT * from products WHERE id = ?"},
	}

	queryRowerStub := func(success bool) *QueryRowerMock {
		return &QueryRowerMock{
			QueryRowFunc: func(_ string, _ ...interface{}) *sql.Row {
				if !success {
					return nil
				}
				return &sql.Row{}
			},
		}
	}

	tests := []struct {
		name           string
		square         SquareSql
		queryRowerStub func(success bool) *QueryRowerMock
		prCallsNumber  int
		psArg          bool
		prepareArg     string
		wantErr        bool
	}{
		{
			name:           "query not found",
			square:         square,
			queryRowerStub: queryRowerStub,
			psArg:          true,
			wantErr:        true,
			prepareArg:     "insert",
			prCallsNumber:  0,
		},
		{
			name:           "error returned by db",
			square:         square,
			queryRowerStub: queryRowerStub,
			psArg:          false,
			wantErr:        true,
			prepareArg:     "select",
			prCallsNumber:  1,
		},
		{
			name:           "successful query preparation",
			square:         square,
			queryRowerStub: queryRowerStub,
			psArg:          true,
			wantErr:        false,
			prepareArg:     "select",
			prCallsNumber:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.queryRowerStub(tt.psArg)
			stmt, err := tt.square.QueryRow(p, tt.prepareArg, "1")
			prCallsNumber := p.CallNumber()
			if tt.wantErr {
				assert.Nil(t, stmt)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stmt)
			}

			assert.Equal(t, prCallsNumber, tt.prCallsNumber)
		})
	}
}

func TestQueryRowContext(t *testing.T) {
	square := SquareSql{
		queries: map[string]string{"select": "SELECT * from products WHERE id = ?"},
	}

	queryRowerContextStub := func(success bool) *QueryRowerContextMock {
		return &QueryRowerContextMock{
			QueryRowContextFunc: func(_ context.Context, _ string, _ ...interface{}) *sql.Row {
				if !success {
					return nil
				}
				return &sql.Row{}
			},
		}
	}

	ctx := context.Background()

	tests := []struct {
		name                  string
		square                SquareSql
		queryRowerContextStub func(success bool) *QueryRowerContextMock
		prCallsNumber         int
		psArg                 bool
		prepareArg            string
		wantErr               bool
	}{
		{
			name:                  "query not found",
			square:                square,
			queryRowerContextStub: queryRowerContextStub,
			psArg:                 true,
			wantErr:               true,
			prepareArg:            "insert",
			prCallsNumber:         0,
		},
		{
			name:                  "error returned by db",
			square:                square,
			queryRowerContextStub: queryRowerContextStub,
			psArg:                 false,
			wantErr:               true,
			prepareArg:            "select",
			prCallsNumber:         1,
		},
		{
			name:                  "successful query preparation",
			square:                square,
			queryRowerContextStub: queryRowerContextStub,
			psArg:                 true,
			wantErr:               false,
			prepareArg:            "select",
			prCallsNumber:         1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.queryRowerContextStub(tt.psArg)
			stmt, err := tt.square.QueryRowContext(ctx, p, tt.prepareArg, "1")
			prCallsNumber := p.CallNumber()
			if tt.wantErr {
				assert.Nil(t, stmt)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stmt)
			}

			assert.Equal(t, prCallsNumber, tt.prCallsNumber)
		})
	}
}

type result struct{}

func (r result) LastInsertId() (int64, error) {
	return 1, nil
}

func (r result) RowsAffected() (int64, error) {
	return 1, nil
}

func TestExec(t *testing.T) {
	square := SquareSql{
		queries: map[string]string{"select": "SELECT * from products WHERE id = ?"},
	}

	execerStub := func(err error) *ExecerMock {
		return &ExecerMock{
			ExecFunc: func(_ string, _ ...interface{}) (sql.Result, error) {
				if err != nil {
					return nil, err
				}
				return result{}, nil
			},
		}
	}

	tests := []struct {
		name          string
		square        SquareSql
		execerStub    func(err error) *ExecerMock
		prCallsNumber int
		psArg         error
		prepareArg    string
		wantErr       bool
	}{
		{
			name:          "error returned by db",
			square:        square,
			execerStub:    execerStub,
			psArg:         errors.New("critical error"),
			wantErr:       true,
			prepareArg:    "select",
			prCallsNumber: 1,
		},
		{
			name:          "not found",
			square:        square,
			execerStub:    execerStub,
			psArg:         nil,
			wantErr:       true,
			prepareArg:    "insert",
			prCallsNumber: 0,
		},
		{
			name:          "successful query preparation",
			square:        square,
			execerStub:    execerStub,
			psArg:         nil,
			wantErr:       false,
			prepareArg:    "select",
			prCallsNumber: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.execerStub(tt.psArg)
			stmt, err := tt.square.Exec(p, tt.prepareArg, "1")
			prCallsNumber := p.CallNumber()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, stmt)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stmt)
			}

			assert.Equal(t, prCallsNumber, tt.prCallsNumber)
		})
	}
}

func TestExecContext(t *testing.T) {
	square := SquareSql{
		queries: map[string]string{"select": "SELECT * from products WHERE id = ?"},
	}

	execerContextStub := func(err error) *ExecerContextMock {
		return &ExecerContextMock{
			ExecContextFunc: func(_ context.Context, _ string, _ ...interface{}) (sql.Result, error) {
				if err != nil {
					return nil, err
				}
				return result{}, nil
			},
		}
	}

	ctx := context.Background()

	tests := []struct {
		name              string
		square            SquareSql
		execerContextStub func(err error) *ExecerContextMock
		prCallsNumber     int
		psArg             error
		prepareArg        string
		wantErr           bool
	}{
		{
			name:              "error returned by db",
			square:            square,
			execerContextStub: execerContextStub,
			psArg:             errors.New("critical error"),
			wantErr:           true,
			prepareArg:        "select",
			prCallsNumber:     1,
		},
		{
			name:              "not found",
			square:            square,
			execerContextStub: execerContextStub,
			psArg:             nil,
			wantErr:           true,
			prepareArg:        "insert",
			prCallsNumber:     0,
		},
		{
			name:              "successful query preparation",
			square:            square,
			execerContextStub: execerContextStub,
			psArg:             nil,
			wantErr:           false,
			prepareArg:        "select",
			prCallsNumber:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.execerContextStub(tt.psArg)
			stmt, err := tt.square.ExecContext(ctx, p, tt.prepareArg, "1")
			prCallsNumber := p.CallNumber()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, stmt)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stmt)
			}

			assert.Equal(t, prCallsNumber, tt.prCallsNumber)
		})
	}
}

func TestLoad(t *testing.T) {
	sqlFile := `
	-- name: all-products
	select * from products
	`

	expectedQuery := "select * from products"

	q, err := Load(strings.NewReader(sqlFile))
	assert.NoError(t, err)
	assert.Equal(t, q.queries["all-products"], expectedQuery)

	raw, err := q.Raw("all-products")
	assert.NoError(t, err)
	assert.Equal(t, raw, expectedQuery)
}

func TestQueries(t *testing.T) {
	expectedQueryMap := map[string]string{
		"select": "SELECT * from products",
		"insert": "INSERT INTO products (?, ?, ?)",
	}

	q, err := LoadFromString(`
	-- name: select
	SELECT * from products
	-- name: insert
	INSERT INTO products (?, ?, ?)
	`)
	assert.NoError(t, err)

	got := q.QueryMap()
	assert.Equal(t, got, expectedQueryMap)
}

func TestMergeHaveBothQueries(t *testing.T) {
	expectedQueryMap := map[string]string{
		"query-a": "SELECT * FROM a",
		"query-b": "SELECT * FROM b",
	}

	a, err := LoadFromString("--name: query-a\nSELECT * FROM a")
	assert.NoError(t, err)

	b, err := LoadFromString("--name: query-b\nSELECT * FROM b")
	assert.NoError(t, err)

	c := Merge(a, b)

	got := c.QueryMap()
	assert.Equal(t, got, expectedQueryMap)
}
