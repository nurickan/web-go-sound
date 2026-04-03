import { Transport } from "./transport"
describe("Transport", () => {
  it("starts playing", () => {
    const t = new Transport(); t.start(); expect(t.state).toBe("playing"); t.stop()
  })
  it("stops resets position", () => {
    const t = new Transport(); t.start(); t.stop(); expect(t.position).toBe(0); expect(t.state).toBe("stopped")
  })
})
