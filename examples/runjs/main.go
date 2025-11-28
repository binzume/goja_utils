package main

import (
	"fmt"
	"os"

	"github.com/binzume/goja_utils"
	"github.com/dop251/goja"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: runjs examples/test.js")
		return
	}

	r := goja_utils.NewJsRunnner()

	// Enables Fetch API
	r.Run(func(vm *goja.Runtime) {
		goja_utils.EnableFetch(vm)
	})

	r.Start()
	result, err := r.RunFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v, err2 := r.Await(result)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else if v != goja.Undefined() {
		fmt.Println(v)
	}

	r.Wait()
}
