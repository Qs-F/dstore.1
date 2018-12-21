package dstore

import (
	"context"
	"database/sql"
	"testing"
)

type RDB struct {
	Path string
	DB   *sql.DB
}

func New(path string) *RDB {
	return &RDB{
		Path: path,
		DB:   nil,
	}
}

func (rdb *RDB) Open(ctx context.Context) (Storer, error) {
	db, err := sql.Open("sqlite", rdb.Path)
	if err != nil {
		return nil, err
	}
	rdb.DB = db
	return rdb, nil
}

func (rdb *RDB) Close(ctx context.Context) (err error) {
	return rdb.DB.Close()
}

func (rdb *RDB) Get(ctx context.Context, key Key, lock chan struct{}) (value chan Value, err error) {
	return nil, nil
}

func (rdb *RDB) Put(ctx context.Context, key Key, value Value, lock chan struct{}) (err error) {
	return nil
}
func (rdb *RDB) Remove(ctx context.Context, key Key, lock chan struct{}) (err error) {
	return nil
}

func (rdb *RDB) Search(ctx context.Context, search Searcher, lock chan struct{}) (err error) {
	return nil
}

func (rdb *RDB) Ping(ctx context.Context) (err error) {
	return rdb.DB.PingContext(ctx)
}

func TestRDB(t *testing.T) {
}

func TestKVS(t *testing.T) {
}
