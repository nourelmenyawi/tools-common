package hofp

import (
	"github.com/couchbase/tools-common/maths"
	"github.com/couchbase/tools-common/system"
)

// Options encapsulates the available options which can be used when creating a worker pool.
type Options struct {
	// Size dictates the number of goroutines created to process incoming functions. Defaults to the number of vCPUs.
	Size int

	// BufferMultiplier is the multiplier used when determining how may functions can be buffered for processioning
	// before calls to 'Queue' block. This value is multiplied by the number of goroutines, and defaults to one.
	BufferMultiplier int

	// LogPrefix is the prefix used when logging errors which occur once teardown has already begun. Defaults to
	// '(hofp)'.
	LogPrefix string
}

func (o *Options) defaults() {
	if o.Size == 0 {
		o.Size = system.NumCPU()
	}

	o.BufferMultiplier = maths.Max(1, o.BufferMultiplier)

	if o.LogPrefix == "" {
		o.LogPrefix = "(hofp)"
	}
}
