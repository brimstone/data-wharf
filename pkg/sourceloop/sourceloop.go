package sourceloop

import (
	"time"

	wharf "github.com/brimstone/data-wharf"
)

type Options struct {
	Lake wharf.Lake
}

type sourceloop struct {
	lake   wharf.Lake
	ticker *time.Ticker
}

func New(options ...*Options) (*sourceloop, error) {
	s := &sourceloop{}
	for _, o := range options {
		if o.Lake != nil {
			s.lake = o.Lake
		}
	}
	return s, nil
}

func (s *sourceloop) Start() {
	s.ticker = time.NewTicker(time.Second)
	go s.run()
}

func (s *sourceloop) run() {
	i := 0
	for {
		select {
		case <-s.ticker.C:
			s.lake.Insert("sourceloop", 1, i)
		}
	}
}
