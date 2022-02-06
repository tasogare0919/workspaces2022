package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Return from New shoud not be nil")
	}

	trcer.Trace("Hello tracer package.")
	if buf.String() != "Hello trace package.\n" {
		t.Errorf("Trace shoud not write '%s'.", buf.String())
	}
}

func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Trace("something")
}