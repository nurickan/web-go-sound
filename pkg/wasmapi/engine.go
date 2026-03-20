package wasmapi; import ("encoding/json"; "github.com/nurickan/web-go-sound/internal/dsp"; "github.com/nurickan/web-go-sound/internal/synth"; "github.com/nurickan/web-go-sound/internal/patch")
type WASMSession struct { engine *synth.Engine; cp patch.Patch }
func NewWASMSession(poly int) *WASMSession { return &WASMSession{engine: synth.NewEngine(dsp.SampleRate, poly), cp: patch.DefaultPatch()} }
func (s *WASMSession) Handle(id int64, raw json.RawMessage) Response {
	var req Request; json.Unmarshal(raw, &req)
	switch req.Method {
	case "noteOn": return s.handleNoteOn(req)
	case "noteOff": return s.handleNoteOff(req)
	case "allNotesOff": s.engine.AllNotesOff(); return Success(id, nil)
	default: return Error(id, 404, "unknown: "+req.Method)
	}
}
func (s *WASMSession) handleNoteOn(req Request) Response {
	var p struct { Note int; Velocity float64 }; json.Unmarshal(req.Params, &p)
	s.engine.NoteOn(p.Note, p.Velocity, s.cp); return Success(req.ID, nil)
}
func (s *WASMSession) handleNoteOff(req Request) Response {
	var p struct { Note int }; json.Unmarshal(req.Params, &p)
	s.engine.NoteOff(p.Note); return Success(req.ID, nil)
}
