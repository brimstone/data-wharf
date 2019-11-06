package wharf_test

import (
	"testing"
	"time"

	"github.com/brimstone/data-wharf/pkg/lakemem"
	"github.com/brimstone/data-wharf/pkg/martcounter"
	"github.com/brimstone/data-wharf/pkg/sourceloop"
	"github.com/brimstone/data-wharf/pkg/warehousemem"
)

func Test_Simple(t *testing.T) {

	// Create the lake to hold the raw events from the source
	l, err := lakemem.New()
	if err != nil {
		panic(err)
	}

	// Create a simple source that just runs a loop making numbers
	s, err := sourceloop.New(&sourceloop.Options{
		Lake: l,
	})
	if err != nil {
		panic(err)
	}

	// Start the source collecting
	s.Start()

	// Create a warehouse and connect it to the lake
	w, err := warehousemem.New(&warehousemem.Options{
		Lake: l,
	})

	// Create a mart to show off the data
	m, err := martcounter.New(&martcounter.Options{
		Warehouse: w,
	})
	time.Sleep(time.Second * 3)
	// The methods on a mart are freeformed for now
	m.Show()
}
