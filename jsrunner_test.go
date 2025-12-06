package goja_utils

import (
	"testing"
)

func TestJsRunner(t *testing.T) {
	r := NewJsRunnner()

	r.Start()

	r.Wait()
}
