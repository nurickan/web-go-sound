package main; import ("syscall/js"; "github.com/nurickan/web-go-sound/pkg/wasmapi")
var session *wasmapi.WASMSession
func main() {
	c := make(chan struct{}, 0); session = wasmapi.NewWASMSession(16)
	js.Global().Set("synthHandle", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer func() { if r := recover(); r != nil { js.Global().Call("console.error", "panic:", r) } }()
		if len(args) < 1 { return map[string]interface{}{"error": map[string]interface{}{"code": 400, "message": "missing"}} }
		return session.Handle(0, []byte(args[0].String()))
	}))
	js.Global().Set("synthReady", js.ValueOf(true)); <-c
}
