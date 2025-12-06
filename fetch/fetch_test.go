package fetch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/binzume/goja_utils"
	"github.com/dop251/goja"
)

func newTestServer() *httptest.Server {
	var sampleHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "Not Found.", http.StatusNotFound)
			return
		}

		fmt.Fprintf(w, `{"status": "ok"}`)
	})
	return httptest.NewServer(sampleHandler)
}

func TestFetch(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	r := goja_utils.NewJsRunnner()

	// Enables Fetch API
	r.Run(func(vm *goja.Runtime) {
		Enable(vm)
		vm.GlobalObject().Set("testServerUrl", ts.URL)
	})

	r.Start()

	ret, err := r.RunFile("fetch_test.js")
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
