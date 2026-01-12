export interface AppState {
  wasmReady: boolean;
  audioReady: boolean;
  theme: "dark" | "light";
  setWasmReady: (v: boolean) => void;
  setAudioReady: (v: boolean) => void;
  setTheme: (t: "dark" | "light") => void;
}

export function createAppState(): AppState {
  let wasmReady = false;
  let audioReady = false;
  let theme: "dark" | "light" = "dark";
  const listeners: Set<() => void> = new Set();

  function notify() { listeners.forEach(l => l()); }

  return {
    get wasmReady() { return wasmReady; },
    get audioReady() { return audioReady; },
    get theme() { return theme; },
    setWasmReady(v: boolean) { wasmReady = v; notify(); },
    setAudioReady(v: boolean) { audioReady = v; notify(); },
    setTheme(t: "dark" | "light") { theme = t; notify(); },
  };
}
