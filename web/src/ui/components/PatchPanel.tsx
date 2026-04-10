import React from "react"
interface PatchPanelProps { patch: any; onPatchChange: (p: any) => void }
export const PatchPanel: React.FC<PatchPanelProps> = ({ patch, onPatchChange }) => (
  <div className="patch-panel">
    <h3>{patch.name}</h3>
    <div className="oscillators">{patch.oscillators?.map((osc: any, i: number) => (
      <div key={i} className="osc-section">
        <span>OSC {i + 1}</span>
        <select value={osc.waveform} onChange={e => { const o = [...patch.oscillators]; o[i] = { ...osc, waveform: parseInt(e.target.value) }; onPatchChange({ ...patch, oscillators: o }) }}>
          <option value={0}>Sine</option><option value={1}>Saw</option><option value={2}>Square</option><option value={3}>Noise</option>
        </select>
      </div>
    ))}</div>
  </div>
)
