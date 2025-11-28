package goja_utils

import (
	"strings"
	"testing"

	"github.com/dop251/goja"
)

func TestReadFile(t *testing.T) {
	const TestScript = `
const fs = require("fs");
fs.readFileSync("testdata/test.txt");
`

	r := NewJsRunnner()

	result, err := r.RunScript("test", TestScript)
	if err != nil {
		t.Fatalf("Failed to load %v", err)
	}

	if str, ok := result.Export().(string); ok {
		if strings.TrimSpace(str) != "Hello" {
			t.Errorf("unexpected result %v", result)
		} else {
			t.Log(str)
		}
	} else {
		t.Errorf("not a string: %v", result)
	}
}

func TestReadFile_BUffer(t *testing.T) {
	const TestScript = `
const fs = require("fs");
fs.readFileSync("testdata/test.txt", {encoding: 'buffer'});
`

	r := NewJsRunnner()

	result, err := r.RunScript("test", TestScript)
	if err != nil {
		t.Fatalf("Failed to load %v", err)
	}

	if buf, ok := result.Export().(goja.ArrayBuffer); ok {
		if strings.TrimSpace(string(buf.Bytes())) != "Hello" {
			t.Errorf("unexpected result %v", result)
		} else {
			t.Log(buf)
		}
	} else {
		t.Errorf("not a string: %v", result)
	}
}
