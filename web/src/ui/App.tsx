import React from "react"; import { SynthView } from "./views/SynthView"; import { PatchView } from "./views/PatchView"
export const App: React.FC = () => (<div className="app"><header><h1>Web-Go-Sound</h1></header><main><SynthView /><PatchView /></main></div>)
