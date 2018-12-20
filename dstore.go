package dstore

import "context"

type Key interface {
	String() string
}

type Value interface {
	Bytes() []byte
}

type Data struct {
	Key   Key
	Value Value
}

type Searcher interface {
	Search() []Data
}

type Storer interface {
	Get(ctx context.Context, key Key, lock chan struct{}) (value chan Value, err error)
	Put(ctx context.Context, key Key, value Value, lock chan struct{}) (err error)
	Remove(ctx context.Context, key Key, lock chan struct{}) (err error)
	Search(ctx context.Context, search Searcher, lock chan struct{}) (err error)
	Open(ctx context.Context) (Store Storer, err error)
	Ping(ctx context.Context) (err error)
	Close(ctx context.Context) (err error)
}
