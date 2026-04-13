import { useCallback } from "react"; import { useDispatch, useSelector } from "react-redux"; import { RootState, noteOn, noteOff } from "../store"
export function useSynth() {
  const activeNotes = useSelector((s: RootState) => s.synth.activeNotes); const dispatch = useDispatch()
  const play = useCallback((n: number) => dispatch(noteOn(n)), [dispatch])
  const release = useCallback((n: number) => dispatch(noteOff(n)), [dispatch])
  return { activeNotes, play, release }
}
