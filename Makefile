.PHONY: setup dev build test lint clean build-wasm verify-release

setup:
	cd web && npm install

dev:
	cd web && npm run dev

build:
	cd web && npm run build

test:
	cd web && npm test
	go test ./...

lint:
	cd web && npm run lint
	golangci-lint run ./...

clean:
	rm -rf web/dist web/node_modules

build-wasm:
	GOOS=js GOARCH=wasm go build -o web/public/wasm/synth.wasm ./cmd/wasm
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" web/public/wasm/

verify-release:
	go test ./... -count=1
	cd web && npm run build
