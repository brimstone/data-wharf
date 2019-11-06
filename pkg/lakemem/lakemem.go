package lakemem

import (
	"sync"

	wharf "github.com/brimstone/data-wharf"
)

type Options struct {
}

type lakemem struct {
	mem       sync.Map
	callbacks []wharf.LakeCallback
}

type mapKey struct {
	Key     string
	Version int
}

func New(options ...*Options) (*lakemem, error) {
	l := &lakemem{}
	return l, nil
}

func (l *lakemem) AddCallback(c wharf.LakeCallback) (cid string, err error) {
	l.callbacks = append(l.callbacks, c)
	return "", nil
}

func (l *lakemem) Insert(key string, version int, payload interface{}) (err error) {
	k := mapKey{
		Key:     key,
		Version: version,
	}
	l.mem.Store(k, payload)
	for _, c := range l.callbacks {
		go c(key, version, payload)
	}
	return nil
}
