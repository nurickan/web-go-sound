import { SynthBridge } from "../wasm/bridge"
export class SynthEngine {
  private bridge: SynthBridge
  constructor() { this.bridge = new SynthBridge() }
  async init(): Promise<void> { await this.bridge.init() }
  noteOn(n: number, v: number = 0.8): void { this.bridge.noteOn(n, v) }
  noteOff(n: number): void { this.bridge.noteOff(n) }
  allNotesOff(): void { this.bridge.allNotesOff() }
}
