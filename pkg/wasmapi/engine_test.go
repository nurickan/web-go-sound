package wasmapi; import ("encoding/json"; "testing")
func TestWASMSessionNoteOn(t *testing.T) {
	s := NewWASMSession(8); req := Request{Method: "noteOn", ID: 1, Params: json.RawMessage(`{"note":69,"velocity":0.8}`)}
	data, _ := json.Marshal(req); r := s.Handle(1, data)
	if r.Error != nil { t.Errorf("err: %s", r.Error.Message) }
}
func TestWASMSessionAllOff(t *testing.T) {
	s := NewWASMSession(8); req := Request{Method: "allNotesOff", ID: 2}; data, _ := json.Marshal(req)
	if r := s.Handle(2, data); r.Error != nil { t.Errorf("err: %s", r.Error.Message) }
}
