export type TransportState = "stopped" | "playing" | "paused"
export class Transport {
  state: TransportState = "stopped"; bpm = 120; position = 0
  onTick: ((pos: number) => void) | null = null
  private id = 0; private interval = 60000 / (120 * 24)
  start(): void {
    if (this.state === "playing") return; this.state = "playing"
    const tick = () => { if (this.state !== "playing") return; this.position++; this.onTick?.(this.position); this.id = window.setTimeout(tick, this.interval) }
    this.id = window.setTimeout(tick, this.interval)
  }
  stop(): void { this.state = "stopped"; window.clearTimeout(this.id); this.position = 0 }
  pause(): void { this.state = "paused"; window.clearTimeout(this.id) }
  setBPM(b: number): void { this.bpm = Math.max(20, Math.min(300, b)); this.interval = 60000 / (this.bpm * 24) }
}
