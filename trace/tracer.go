package trace

import (
	"fmt"
	"io"
)

type Tracer struct {
	out io.Writer
}

func New(w io.Writer) Tracer {
	return Tracer{out: w}
}

func (t *Tracer) Trace(a ...interface{}) {
	if t.out == nil {
		return
	}
	fmt.Fprintln(t.out, a...)
}
