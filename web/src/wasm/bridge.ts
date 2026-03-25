/// <reference lib="dom" />
export interface WASMResponse { result?: unknown; error?: { code: number; message: string } }
export class SynthBridge {
  private ready = false;
  async init(): Promise<void> {
    if (this.ready) return;
    await new Promise<void>((resolve) => {
      const check = () => { if ((window as any).synthReady) { this.ready = true; resolve() } else { setTimeout(check, 50) } }
      check()
    })
  }
  noteOn(note: number, velocity: number): void {
    (window as any).synthHandle(JSON.stringify({ method: "noteOn", params: { note, velocity } }))
  }
  noteOff(note: number): void {
    (window as any).synthHandle(JSON.stringify({ method: "noteOff", params: { note } }))
  }
  allNotesOff(): void { (window as any).synthHandle(JSON.stringify({ method: "allNotesOff" })) }
}
