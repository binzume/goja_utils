package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/binzume/goja_utils"
	"github.com/dop251/goja"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

var HttpClient Client = http.DefaultClient

type JsMap map[string]any

func (o JsMap) GetString(name, def string) string {
	if o == nil {
		return def
	}
	if v, ok := o[name]; ok {
		return fmt.Sprint(v)
	}
	return def
}

func makeFetch(vm *goja.Runtime) any {
	r := goja_utils.GetTaskQueue(vm)
	return func(url string, options JsMap) any {
		method := options.GetString("method", "GET")
		var body io.Reader
		if _, ok := options["body"]; ok {
			body = strings.NewReader(options.GetString("body", ""))
		}
		req, err := http.NewRequest(method, url, body)
		if h, ok := options["headers"].(map[string]any); ok {
			for k, v := range h {
				req.Header.Add(k, fmt.Sprint(v))
			}
		}
		return r.StartGoroutineTask(func() (any, error) {
			if err != nil {
				return nil, err
			}

			resp, err := HttpClient.Do(req)
			if err != nil {
				return nil, err
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resHeaders := map[string]any{}
			for k, v := range resp.Header {
				resHeaders[k] = v
			}
			resp.Location()
			return map[string]any{
				"ok":         resp.StatusCode >= 200 && resp.StatusCode < 300,
				"status":     resp.StatusCode,
				"statusText": resp.Status,
				"url":        resp.Request.URL.String(),
				"headers":    resHeaders,
				"text":       func() string { return string(body) },
				"json": func() (any, error) {
					var data any
					err := json.Unmarshal(body, &data)
					return data, err
				},
				"arrayBuffer": func() any { return vm.NewArrayBuffer(body) },
				"bytes": func() (any, error) {
					return vm.New(vm.Get("Uint8Array"), vm.ToValue(vm.NewArrayBuffer(body)))
				},
			}, nil
		})
	}
}

func Enable(vm *goja.Runtime) {
	vm.Set("fetch", makeFetch(vm))
}
