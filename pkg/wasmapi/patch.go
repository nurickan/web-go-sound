package wasmapi; import ("encoding/json"; "github.com/nurickan/web-go-sound/internal/patch")
func HandleValidatePatch(id int64, raw json.RawMessage) Response {
	var p patch.Patch; if err := json.Unmarshal(raw, &p); err != nil { return Error(id, 400, err.Error()) }
	if err := patch.Validate(p); err != nil { return Error(id, 422, err.Error()) }; return Success(id, map[string]string{"status": "valid"})
}
