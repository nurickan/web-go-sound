import React from "react"; import { render, fireEvent } from "@testing-library/react"; import { Knob } from "./Knob"
describe("Knob", () => {
  it("renders label", () => { const { getByText } = render(<Knob label="Cutoff" value={0.5} onChange={() => {}} />); expect(getByText("Cutoff")).toBeDefined() })
  it("calls onChange", () => {
    let v = 0; const { getByRole } = render(<Knob label="V" value={0.5} onChange={x => { v = x }} />)
    fireEvent.change(getByRole("slider"), { target: { value: "0.8" } }); expect(v).toBe(0.8)
  })
})
