package wasmapi; import "encoding/json"
type Request struct { Method string; ID int64; Params json.RawMessage }
type Response struct { ID int64; Result interface{}; Error *ErrorInfo }
type ErrorInfo struct { Code int; Message string }
func Success(id int64, result interface{}) Response { return Response{ID: id, Result: result} }
func Error(id int64, code int, msg string) Response { return Response{ID: id, Error: &ErrorInfo{Code: code, Message: msg}} }
