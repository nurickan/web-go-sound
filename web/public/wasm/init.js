async function initSynth() {
  if (window.synthReady) return;
  const go = new Go();
  const result = await WebAssembly.instantiateStreaming(fetch("wasm/synth.wasm"), go.importObject);
  go.run(result.instance);
}
window.addEventListener("DOMContentLoaded", () => { initSynth(); });
