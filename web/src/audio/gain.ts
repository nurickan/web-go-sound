export class GainNode {
  private ctx: AudioContext; private node: GainNode
  constructor(ctx: AudioContext) { this.ctx = ctx; this.node = ctx.createGain(); this.node.gain.value = 0.5 }
  connect(dest: AudioNode): void { this.node.connect(dest) }
  setVolume(v: number): void { this.node.gain.setTargetAtTime(v, this.ctx.currentTime, 0.02) }
  getVolume(): number { return this.node.gain.value }
}
