package dstore

import (
	"context"
	"database/sql"
	"strconv"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"gopkg.in/gorp.v2"
)

type Table struct {
	ID    int    `db:"id, primarykey"`
	Data1 string `db:"data1"`
	Data2 string `db:"data2"`
}

type RDB struct {
	Path string
	DB   *gorp.DbMap
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
	rdb.DB = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	return rdb, nil
}

func (rdb *RDB) Close(ctx context.Context) (err error) {
	return rdb.DB.Db.Close()
}

type key int

func (key key) String() string {
	return strconv.Itoa(int(key))
}

func (rdb *RDB) Get(ctx context.Context, key Key, lock chan bool) (value chan<- Value, err error) {
	select {
	case <-lock:
		tx, err := rdb.DB.Begin()
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (rdb *RDB) Put(ctx context.Context, key Key, value Value, lock chan bool) (err error) {
	return nil
}
func (rdb *RDB) Remove(ctx context.Context, key Key, lock chan bool) (err error) {
	return nil
}

func (rdb *RDB) Search(ctx context.Context, search Searcher, lock chan bool) (err error) {
	return nil
}

func (rdb *RDB) Ping(ctx context.Context) (err error) {
	return rdb.DB.Db.PingContext(ctx)
}

func TestRDB(t *testing.T) {
}

func TestKVS(t *testing.T) {
}
