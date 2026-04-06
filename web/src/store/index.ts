import { configureStore, createSlice, PayloadAction } from "@reduxjs/toolkit"
interface SynthState { bpm: number; volume: number; activeNotes: number[]; transport: string }
const initial: SynthState = { bpm: 120, volume: 0.5, activeNotes: [], transport: "stopped" }
const slice = createSlice({
  name: "synth", initialState: initial,
  reducers: {
    setBpm(s, a: PayloadAction<number>) { s.bpm = a.payload },
    setVolume(s, a: PayloadAction<number>) { s.volume = a.payload },
    noteOn(s, a: PayloadAction<number>) { if (!s.activeNotes.includes(a.payload)) s.activeNotes.push(a.payload) },
    noteOff(s, a: PayloadAction<number>) { s.activeNotes = s.activeNotes.filter(n => n !== a.payload) },
    setTransport(s, a: PayloadAction<string>) { s.transport = a.payload },
  }
})
export const { setBpm, setVolume, noteOn, noteOff, setTransport } = slice.actions
export const store = configureStore({ reducer: { synth: slice.reducer } })
export type RootState = ReturnType<typeof store.getState>
