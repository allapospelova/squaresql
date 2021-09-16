package squaresql

import (
	"context"
	"database/sql"
)

type PreparerMock struct {
	// PrepareFunc mocks the Prepare method.
	PrepareFunc func(query string) (*sql.Stmt, error)

	// calls tracks calls to the methods.
	calls struct {
		// Prepare holds details about calls to the Prepare method.
		Prepare []struct {
			// Query is the query argument value.
			Query string
		}
	}

	cn int
}

func (mock *PreparerMock) Prepare(query string) (*sql.Stmt, error) {
	mock.cn++
	return mock.PrepareFunc(query)
}

func (mock *PreparerMock) CallNumber() int {
	return mock.cn
}

type PreparerContextMock struct {
	// PrepareContextFunc mocks the PrepareContext method.
	PrepareContextFunc func(ctx context.Context, query string) (*sql.Stmt, error)

	// calls tracks calls to the methods.
	calls struct {
		// PrepareContext holds details about calls to the PrepareContext method.
		PrepareContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
		}
	}

	cn int
}

func (mock *PreparerContextMock) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	mock.cn++
	return mock.PrepareContextFunc(ctx, query)
}

func (mock *PreparerContextMock) CallNumber() int {
	return mock.cn
}

type QueryerMock struct {
	QueryFunc func(_ string, _ ...interface{}) (*sql.Rows, error)
	cn        int
}

func (mock *QueryerMock) Query(name string, args ...interface{}) (*sql.Rows, error) {
	mock.cn++
	return mock.QueryFunc(name, args)
}

func (mock *QueryerMock) CallNumber() int {
	return mock.cn
}

type QueryerContextMock struct {
	QueryContextFunc func(_ context.Context, _ string, _ ...interface{}) (*sql.Rows, error)
	cn               int
}

func (mock *QueryerContextMock) QueryContext(ctx context.Context, name string, args ...interface{}) (*sql.Rows, error) {
	mock.cn++
	return mock.QueryContextFunc(ctx, name, args)
}

func (mock *QueryerContextMock) CallNumber() int {
	return mock.cn
}

type QueryRowerMock struct {
	QueryRowFunc func(_ string, _ ...interface{}) *sql.Row
	cn           int
}

func (mock *QueryRowerMock) QueryRow(name string, args ...interface{}) *sql.Row {
	mock.cn++
	return mock.QueryRowFunc(name, args)
}

func (mock *QueryRowerMock) CallNumber() int {
	return mock.cn
}

type QueryRowerContextMock struct {
	QueryRowContextFunc func(_ context.Context, _ string, _ ...interface{}) *sql.Row
	cn                  int
}

func (mock *QueryRowerContextMock) QueryRowContext(ctx context.Context, name string, args ...interface{}) *sql.Row {
	mock.cn++
	return mock.QueryRowContextFunc(ctx, name, args)
}

func (mock *QueryRowerContextMock) CallNumber() int {
	return mock.cn
}

type ExecerMock struct {
	ExecFunc func(_ string, _ ...interface{}) (sql.Result, error)
	cn       int
}

func (mock *ExecerMock) Exec(name string, args ...interface{}) (sql.Result, error) {
	mock.cn++
	return mock.ExecFunc(name, args)
}

func (mock *ExecerMock) CallNumber() int {
	return mock.cn
}

type ExecerContextMock struct {
	ExecContextFunc func(_ context.Context, _ string, _ ...interface{}) (sql.Result, error)
	cn              int
}

func (mock *ExecerContextMock) ExecContext(ctx context.Context, name string, args ...interface{}) (sql.Result, error) {
	mock.cn++
	return mock.ExecContextFunc(ctx, name, args)
}

func (mock *ExecerContextMock) CallNumber() int {
	return mock.cn
}
