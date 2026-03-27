let ctx: AudioContext | null = null
export function getAudioContext(): AudioContext {
  if (!ctx) ctx = new AudioContext()
  return ctx
}
export function resumeAudio(): Promise<void> {
  const c = getAudioContext()
  return c.state === "suspended" ? c.resume() : Promise.resolve()
}
