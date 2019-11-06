package warehousemem

import (
	"errors"
	"sync"

	wharf "github.com/brimstone/data-wharf"
	"github.com/brimstone/logger"
)

type Options struct {
	Lake wharf.Lake
}

type warehousemem struct {
	lake wharf.Lake
	jobs []wharf.Job
	lock sync.Mutex
	kv   map[string]interface{}
}

func New(options ...*Options) (*warehousemem, error) {
	w := &warehousemem{
		kv: make(map[string]interface{}),
	}
	for _, o := range options {
		if o.Lake != nil {
			w.lake = o.Lake
			w.Plumb(o.Lake)
		}
	}
	return w, nil
}

func (w *warehousemem) Plumb(lake wharf.Lake) (err error) {
	w.lake.AddCallback(w.callback)
	return nil
}

func (w *warehousemem) callback(key string, version int, payload interface{}) {
	log := logger.New()
	log.Debug("Got the callback!",
		log.Field("key", key),
		log.Field("version", version),
		log.Field("payload", payload),
	)
	for _, c := range w.jobs {
		c.Callback(w, key, version, payload)
	}
}

func (w *warehousemem) AddJob(job wharf.Job) {
	w.jobs = append(w.jobs, job)
	if job.Backfill() {
		// TODO go w.backfill(job)
	}
}

func (w *warehousemem) GetOne(key string) (interface{}, error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	p, ok := w.kv[key]
	if !ok {
		return nil, errors.New("Not found")
	}
	return p, nil
}

func (w *warehousemem) Inc(key string) error {
	w.lock.Lock()
	defer w.lock.Unlock()
	u, ok := w.kv[key]
	var v uint64
	if ok {
		v = u.(uint64)
	}
	v++
	w.kv[key] = v
	return nil
}
