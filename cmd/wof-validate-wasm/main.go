//go:build wasmjs
package main

import (
	"log"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-validate"
	"github.com/whosonfirst/go-whosonfirst-validate/wasm"	
)

func main() {

	opts := validate.DefaultValidateOptions()
	validate_func := wasm.ValidateFunc(opts)

	defer validate_func.Release()

	js.Global().Set("wof_validate", validate_func)

	c := make(chan struct{}, 0)

	log.Println("Who's On First validate_feature WASM binary initialized")
	<-c
}
