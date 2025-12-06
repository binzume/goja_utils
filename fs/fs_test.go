package goja_utils

import (
	"strings"
	"testing"

	"github.com/binzume/goja_utils"
	"github.com/dop251/goja"
)

func TestReadFile(t *testing.T) {
	r := goja_utils.NewJsRunnner()

	r.Start()

	ret, err := r.RunFile("fs_test.js")
	if err != nil {
		t.Fatalf("Failed to load %v", err)
	}

	resultValue, err2 := r.Await(ret)
	if err2 != nil {
		t.Fatalf("Failed to test: %v", err2)
	}

	result := resultValue.Export()
	if str, ok := result.(string); ok && strings.TrimSpace(str) == "pass" {
		t.Log(str)
	} else {
		t.Errorf("Unexpected result: %v", result)
	}
}

func TestReadFile_BUffer(t *testing.T) {
	const TestScript = `
const fs = require("fs");
fs.readFileSync("../testdata/test.txt", {encoding: 'buffer'});
`

	r := goja_utils.NewJsRunnner()

	result, err := r.RunScript("test", TestScript)
	if err != nil {
		t.Fatalf("Failed to load %v", err)
	}

	if buf, ok := result.Export().(goja.ArrayBuffer); ok {
		if strings.TrimSpace(string(buf.Bytes())) != "Hello" {
			t.Errorf("Unexpected result %v", result)
		} else {
			t.Log(buf)
		}
	} else {
		t.Errorf("not a string: %#v", result)
	}
}
