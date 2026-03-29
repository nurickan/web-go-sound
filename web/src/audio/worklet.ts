export class SynthWorklet {
  private node: AudioWorkletNode | null = null
  async init(ctx: AudioContext): Promise<void> {
    await ctx.audioWorklet.addModule("wasm/worklet-processor.js")
    this.node = new AudioWorkletNode(ctx, "synth-worklet")
    this.node.connect(ctx.destination)
  }
  postMessage(msg: unknown): void { this.node?.port.postMessage(msg) }
  onMessage(cb: (msg: unknown) => void): void {
    if (this.node) this.node.port.onmessage = (e) => cb(e.data)
  }
}
