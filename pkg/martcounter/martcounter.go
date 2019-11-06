package martcounter

import (
	"fmt"

	wharf "github.com/brimstone/data-wharf"
)

type Options struct {
	Warehouse wharf.Warehouse
}

type martcounter struct {
	warehouse wharf.Warehouse
}

func New(o ...*Options) (*martcounter, error) {
	m := &martcounter{}
	return m, nil
}

func (m *martcounter) Show() {
	fmt.Println("Mart counter!")
}
