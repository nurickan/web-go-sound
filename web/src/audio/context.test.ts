import { getAudioContext } from "./context"
describe("AudioContext", () => {
  it("returns singleton", () => {
    const a = getAudioContext()
    const b = getAudioContext()
    expect(a).toBe(b)
  })
})
