package goja_utils

import (
	"io"
	"io/fs"
	"os"

	"github.com/binzume/goja_utils"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func statSync(name string, options map[string]any) (any, error) {
	s, err := os.Stat(name)
	if err != nil {
		if t, ok := options["throwIfNoEntry"].(bool); ok && t {
			return nil, err
		}
		return nil, nil
	}
	return map[string]any{
		"size":        s.Size(),
		"mtimeMs":     s.ModTime().UnixMilli(),
		"isDirectory": func() bool { return s.IsDir() },
	}, nil
}

func rmSync(name string, options map[string]any) error {
	if t, ok := options["recursive"].(bool); ok && t {
		return os.RemoveAll(name)
	}

	return os.Remove(name)
}

func mkdirSync(name string, options map[string]any) error {
	mode := 0777
	if m, ok := options["mode"].(int); ok {
		mode = m
	}
	return os.Mkdir(name, fs.FileMode(mode))
}

func renameSync(name, name2 string) error {
	return os.Rename(name, name2)
}

func convOutput(data []byte, vm *goja.Runtime, options map[string]any) any {
	if options != nil {
		switch options["encoding"] {
		case "bytes":
			return data
		case "buffer":
			return vm.NewArrayBuffer(data)
		}
	}

	return string(data)
}

func makeReadFileSync(vm *goja.Runtime) any {
	return func(path string, options map[string]any) (any, error) {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		data, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		return convOutput(data, vm, options), nil
	}
}

func WriteFileSync(path, text string) error {
	f, err := os.Create(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	_, err = f.WriteString(text)
	return err
}

func AppendFileSync(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil
	}
	defer f.Close()
	_, err = f.WriteString(text)
	return err
}

func ReadFileAsync(r goja_utils.TaskQueue) any {
	return func(name string) goja.Value {
		return r.StartGoroutineTask(func() (any, error) {
			f, err := os.Open(name)
			if err != nil {
				return nil, err
			}
			defer f.Close()
			data, err := io.ReadAll(f)
			return string(data), err
		})
	}
}

func WriteFileAsync(r goja_utils.TaskQueue) any {
	return func(name, text string) goja.Value {
		return r.StartGoroutineTask(func() (any, error) {
			f, err := os.Create(name)
			if err != nil {
				return nil, err
			}
			defer f.Close()
			_, err = f.WriteString(text)
			return nil, err
		})
	}
}

func SetupFsPromises(runtime *goja.Runtime, o *goja.Object) {
	if r := goja_utils.GetTaskQueue(runtime); r != nil {
		o.Set("readFile", ReadFileAsync(r))
		o.Set("writeFile", WriteFileAsync(r))
		// TODO: async
		o.Set("stat", statSync)
		o.Set("mkdir", mkdirSync)
		o.Set("rm", rmSync)
		o.Set("unlink", rmSync)
		o.Set("rename", renameSync)
	}
}

func RequireFs(runtime *goja.Runtime, module *goja.Object) {
	o := module.Get("exports").(*goja.Object)
	o.Set("readFileSync", makeReadFileSync(runtime))
	o.Set("appendFileSync", AppendFileSync)
	o.Set("writeFileSync", WriteFileSync)
	o.Set("statSync", statSync)
	o.Set("mkdirSync", mkdirSync)
	o.Set("rmSync", rmSync)
	o.Set("unlinkSync", rmSync)
	o.Set("renameSync", renameSync)
	po := runtime.NewObject()
	SetupFsPromises(runtime, po)
	o.Set("promises", po)
}

func EnableFs(runtime *goja.Runtime) {
	runtime.Set("fs", require.Require(runtime, "fs"))
}

func init() {
	require.RegisterCoreModule("fs", RequireFs)
}
