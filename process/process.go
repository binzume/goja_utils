package process

import (
	"io"
	"os"
	"strings"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

type ProcessConfig struct {
	Stdin  io.Reader
	Stdout io.Writer
	Env    map[string]string
}

func NewProcessConfig() *ProcessConfig {
	env := map[string]string{}
	for _, e := range os.Environ() {
		kv := strings.SplitN(e, "=", 2)
		if len(kv) == 2 {
			env[kv[0]] = kv[1]
		}
	}
	return &ProcessConfig{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Env:    env,
	}
}

func createReadable(vm *goja.Runtime, r io.Reader) any {
	if r == nil {
		return nil
	}

	o := vm.NewObject()
	o.Set("read", func(size int, encoding string) any {
		if size == 0 {
			size = 1024
		}
		buf := make([]byte, size)
		n, _ := r.Read(buf)
		if encoding == "bytes" {
			return buf[0:n]
		}
		return string(buf[0:n])
	})
	return o
}

func createWritable(vm *goja.Runtime, w io.Writer) any {
	if w == nil {
		return nil
	}

	o := vm.NewObject()
	o.Set("write", func(data string, encoding string) {
		w.Write([]byte(data))
	})
	return o
}

func setup(vm *goja.Runtime, o *goja.Object, env *ProcessConfig) {
	o.Set("stdin", createReadable(vm, env.Stdin))
	o.Set("stdout", createWritable(vm, env.Stdout))
	o.Set("env", env.Env)
}

// registry.RegisterNativeModule("process", RequireWithJsEnv(env))
func RequireWithConfig(env *ProcessConfig) require.ModuleLoader {
	return func(vm *goja.Runtime, module *goja.Object) {
		setup(vm, module.Get("exports").(*goja.Object), env)
	}
}

func Require(vm *goja.Runtime, module *goja.Object) {
	setup(vm, module.Get("exports").(*goja.Object), NewProcessConfig())
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("process", require.Require(runtime, "process"))
}

func init() {
	require.RegisterCoreModule("process", Require)
}
