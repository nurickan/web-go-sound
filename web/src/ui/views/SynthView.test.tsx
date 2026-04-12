import React from "react"; import { render } from "@testing-library/react"; import { SynthView } from "./SynthView"
describe("SynthView", () => {
  it("renders without crash", () => { const { container } = render(<SynthView />); expect(container).toBeDefined() })
})
