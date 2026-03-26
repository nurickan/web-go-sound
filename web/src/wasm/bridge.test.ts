import { SynthBridge } from "./bridge"
describe("SynthBridge", () => {
  it("exposes methods", () => {
    const b = new SynthBridge()
    expect(b.noteOn).toBeDefined()
    expect(b.noteOff).toBeDefined()
    expect(b.allNotesOff).toBeDefined()
  })
})
