import React, { useState } from "react"; import { PatchPanel } from "../components/PatchPanel"
const DEFAULT_PATCH = { name: "Init", oscillators: [{ waveform: 0, octave: 4, detune: 0, level: 0.8, pan: 0 }], filter: { type: 0, cutoff: 18000, resonance: 0, enabled: false }, amplitude: { attackTime: 0.01, decayTime: 0.1, sustainLevel: 0.8, releaseTime: 0.3 }, filterEnv: { attackTime: 0.02, decayTime: 0.2, sustainLevel: 0.5, releaseTime: 0.5 }, lfos: [] }
export const PatchView: React.FC = () => {
  const [patch, setPatch] = useState(DEFAULT_PATCH)
  return (<div className="patch-view"><h2>Patch Editor</h2><PatchPanel patch={patch} onPatchChange={setPatch} /></div>)
}
