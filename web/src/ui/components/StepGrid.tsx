import React from "react"
interface StepGridProps { steps: boolean[]; onToggle: (i: number) => void }
export const StepGrid: React.FC<StepGridProps> = ({ steps, onToggle }) => (
  <div className="step-grid">{
    steps.map((a, i) => (
      <button key={i} className={a ? "on" : "off"} onClick={() => onToggle(i)}>{i + 1}</button>
    ))
  }</div>
)
