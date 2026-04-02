import { KeyboardController } from "./keyboard"
describe("KeyboardController", () => {
  it("maps key to midi", () => {
    const k = new KeyboardController(); let note = -1
    k.onNoteOn = (n) => { note = n }; k.handleDown("a"); expect(note).toBe(60)
  })
})
