package warehousemem

import (
	"fmt"

	wharf "github.com/brimstone/data-wharf"
)

type Options struct {
	Lake wharf.Lake
}

type warehousemem struct {
	lake wharf.Lake
}

func New(options ...*Options) (*warehousemem, error) {
	w := &warehousemem{}
	for _, o := range options {
		if o.Lake != nil {
			w.lake = o.Lake
			w.Plumb(o.Lake)
		}
	}
	return w, nil
}

func (w *warehousemem) Plumb(lake wharf.Lake) (err error) {
	w.lake.AddCallback(w.lakecallback)

	return nil
}

func (w *warehousemem) lakecallback(key string, version int, payload interface{}) {
	fmt.Println("Got the callback!")
}
