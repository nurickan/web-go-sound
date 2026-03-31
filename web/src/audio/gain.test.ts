import { GainNode } from "./gain"
describe("GainNode", () => {
  it("sets volume", () => {
    const g = new GainNode(new AudioContext())
    g.setVolume(0.8)
    expect(g.getVolume()).toBe(0.8)
  })
})
