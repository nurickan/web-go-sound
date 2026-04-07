import React from "react"
interface KnobProps { label: string; value: number; min?: number; max?: number; step?: number; onChange: (v: number) => void }
export const Knob: React.FC<KnobProps> = ({ label, value, min = 0, max = 1, step = 0.01, onChange }) => (
  <div className="knob">
    <label>{label}</label>
    <input type="range" min={min} max={max} step={step} value={value}
      onChange={e => onChange(parseFloat(e.target.value))} />
    <span>{value.toFixed(2)}</span>
  </div>
)
