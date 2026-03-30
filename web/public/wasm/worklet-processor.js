class SynthProcessor extends AudioWorkletProcessor {
  constructor() { super(); this.port.onmessage = (e) => { this.port.postMessage({ type: "echo", data: e.data }) } }
  process(_inputs, outputs, _params) {
    const out = outputs[0]; const ch = out[0]
    for (let i = 0; i < ch.length; i++) ch[i] = Math.sin((currentFrame + i) * 0.01) * 0.1
    return true
  }
}
registerProcessor("synth-worklet", SynthProcessor)
