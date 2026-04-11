import React, { useState } from "react"; import { Knob } from "../components/Knob"; import { StepGrid } from "../components/StepGrid"
export const SynthView: React.FC = () => {
  const [steps, setSteps] = useState(Array(16).fill(false)); const [bpm, setBpm] = useState(120); const [vol, setVol] = useState(0.5)
  return (<div className="synth-view">
    <Knob label="BPM" value={bpm} min={20} max={300} step={1} onChange={setBpm} />
    <Knob label="Volume" value={vol} onChange={setVol} />
    <StepGrid steps={steps} onToggle={i => { const s = [...steps]; s[i] = !s[i]; setSteps(s) }} />
  </div>)
}
