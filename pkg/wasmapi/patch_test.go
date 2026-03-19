package wasmapi; import ("encoding/json"; "testing"; "github.com/nurickan/web-go-sound/internal/patch")
func TestValidatePatch(t *testing.T) {
	p := patch.DefaultPatch(); data, _ := json.Marshal(p)
	if r := HandleValidatePatch(1, data); r.Error != nil { t.Errorf("err: %s", r.Error.Message) }
}
func TestValidatePatchInvalid(t *testing.T) {
	p := patch.DefaultPatch(); p.Oscillators = nil; data, _ := json.Marshal(p)
	if r := HandleValidatePatch(2, data); r.Error == nil { t.Error("expected error") }
}
