package wharf

type Source interface {
}

// https://godoc.org/modernc.org/kv
type Lake interface {
	Insert(key string, version int, payload interface{}) (err error)
	AddCallback(Callback) (cid string, err error)
}

type Callback func(key string, version int, payload interface{})

// https://godoc.org/modernc.org/ql
//
type Warehouse interface {
	Plumb(lake Lake) (err error)
	AddJob(Job)

	// FIXME Not sure about this one
	GetOne(key string) (interface{}, error)
	Inc(key string) error
}

type Job interface {
	Backfill() bool
	Callback(w Warehouse, key string, version int, payload interface{}) error
}

type Mart interface {
}
