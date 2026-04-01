export interface KeyMap { [key: string]: number }
const QWERTY_MAP: KeyMap = { a: 60, w: 61, s: 62, e: 63, d: 64, f: 65, t: 66, g: 67, y: 68, h: 69, u: 70, j: 71, k: 72 }
export class KeyboardController {
  private active = new Set<number>(); private map: KeyMap
  constructor(map?: KeyMap) { this.map = map ?? QWERTY_MAP }
  onNoteOn: ((n: number) => void) | null = null; onNoteOff: ((n: number) => void) | null = null
  handleDown(key: string): void {
    const note = this.map[key]; if (note === undefined || this.active.has(note)) return
    this.active.add(note); this.onNoteOn?.(note)
  }
  handleUp(key: string): void {
    const note = this.map[key]; if (note === undefined) return
    this.active.delete(note); this.onNoteOff?.(note)
  }
  allOff(): void { for (const n of this.active) { this.onNoteOff?.(n) }; this.active.clear() }
}
