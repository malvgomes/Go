package arq

import (
	"runtime"
	"testing"
)

func TestDep(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Doesn't work in amd64")
	}
	t.Fail()
}
