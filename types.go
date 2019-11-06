package wharf

type Source interface {
}

// https://godoc.org/modernc.org/kv
type Lake interface {
	Insert(key string, version int, payload interface{}) (err error)
	AddCallback(LakeCallback) (cid string, err error)
}

type LakeCallback func(key string, version int, payload interface{})

// https://godoc.org/modernc.org/ql
//
type Warehouse interface {
	Plumb(lake Lake) (err error)
}

type Mart interface {
}
