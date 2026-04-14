import { useCallback } from "react"; import { useDispatch, useSelector } from "react-redux"; import { RootState, setBpm, setTransport } from "../store"
export function useTransport() {
  const bpm = useSelector((s: RootState) => s.synth.bpm); const transport = useSelector((s: RootState) => s.synth.transport); const dispatch = useDispatch()
  const start = useCallback(() => dispatch(setTransport("playing")), [dispatch])
  const stop = useCallback(() => dispatch(setTransport("stopped")), [dispatch])
  const changeBpm = useCallback((b: number) => dispatch(setBpm(b)), [dispatch])
  return { bpm, transport, start, stop, changeBpm }
}
