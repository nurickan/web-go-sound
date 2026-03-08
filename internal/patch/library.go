package patch; import ("encoding/json"; "os")
type Library struct { presets map[string]Patch }
func NewLibrary() *Library { return &Library{presets: make(map[string]Patch)} }
func (l *Library) LoadFromFile(path string) error {
	data, err := os.ReadFile(path); if err != nil { return err }
	var p Patch; if err := json.Unmarshal(data, &p); err != nil { return err }
	if err := Validate(p); err != nil { return err }; l.presets[p.Name] = p; return nil
}
func (l *Library) Get(name string) (Patch, bool) { p, ok := l.presets[name]; return p, ok }
func (l *Library) List() []string { n := make([]string, 0, len(l.presets)); for k := range l.presets { n = append(n, k) }; return n }
func (l *Library) Add(p Patch) { l.presets[p.Name] = p }
