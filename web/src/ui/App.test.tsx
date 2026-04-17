import React from "react"; import { render } from "@testing-library/react"; import { App } from "./App"
describe("App", () => {
  it("renders header", () => { const { getByText } = render(<App />); expect(getByText("Web-Go-Sound")).toBeDefined() })
})
