package martcounter

import (
	wharf "github.com/brimstone/data-wharf"
	"github.com/brimstone/logger"
)

type Options struct {
	Warehouse wharf.Warehouse
}

type martcounter struct {
	warehouse wharf.Warehouse
}

func New(options ...*Options) (*martcounter, error) {
	m := &martcounter{}
	for _, o := range options {
		if o.Warehouse != nil {
			m.warehouse = o.Warehouse
		}
	}
	return m, nil
}

func (m *martcounter) Show() {
	log := logger.New()
	c, err := m.warehouse.GetOne("sourceloop-count")
	if err != nil {
		panic(err)
	}
	log.Info("Total",
		log.Field("counts", c),
	)
}
