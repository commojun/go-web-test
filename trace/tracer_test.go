package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	tracer.Trace("こんにちは、traceパッケージ")
	if buf.String() != "こんにちは、traceパッケージ\n" {
		t.Errorf("'%s'という誤った文字が出ました", buf.String())
	}
}
