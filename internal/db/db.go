package db

import "context"

type DBAccessor interface {
	FetchRows(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (int, error)
	Exec(ctx context.Context, tag string, result interface{}, query string, args ...interface{}) (interface{}, error)
}
