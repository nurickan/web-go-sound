import { copyFileSync } from "fs";
import { resolve } from "path";

const goroot = process.env.GOROOT || "";
const src = resolve(goroot, "misc/wasm/wasm_exec.js");
const dest = resolve("web/public/wasm/wasm_exec.js");
copyFileSync(src, dest);
console.log("wasm_exec.js copied to", dest);
