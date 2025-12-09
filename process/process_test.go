package process

import (
	"strings"
	"testing"

	"github.com/binzume/goja_utils"
)

func TestProcess(t *testing.T) {
	r := goja_utils.NewJsRunnner()

	config := &ProcessConfig{
		Env: map[string]string{
			"TEST": "Hello",
		},
		Stdin: strings.NewReader("testtest"),
	}
	r.Registry().RegisterNativeModule("process", RequireWithConfig(config))

	r.Start()

	ret, err := r.RunFile("process_test.js")
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
