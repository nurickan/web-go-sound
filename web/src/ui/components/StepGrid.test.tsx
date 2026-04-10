import React from "react"; import { render, fireEvent } from "@testing-library/react"; import { StepGrid } from "./StepGrid"
describe("StepGrid", () => {
  it("renders 16 steps", () => { const { getAllByRole } = render(<StepGrid steps={Array(16).fill(false)} onToggle={() => {}} />); expect(getAllByRole("button").length).toBe(16) })
  it("toggles step", () => {
    let toggled = -1; const { getByText } = render(<StepGrid steps={Array(8).fill(false)} onToggle={i => { toggled = i }} />)
    fireEvent.click(getByText("3")); expect(toggled).toBe(2)
  })
})
