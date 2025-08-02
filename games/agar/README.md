# agar

## Requirements

- Godot: Godot_v4.4.1-stable
  - Godot_v4.4.1-stable_win64.exe

## Structure

- `game` - godot game code
- `gdextensions`
- `server` - go server
- `shared` - shared between game (client) and server

## Info

- Based on [Make an MMO with Godot 4 + Golang by Tristan Batchler](https://www.youtube.com/playlist?list=PLA1tuaTAYPbHAU2ISi_aMjSyZr-Ay7UTJ)
  - https://www.tbat.me/2024/11/09/godot-golang-mmo-part-1
- Repo https://github.com/tristanbatchler/Godot4Go_MMO

## Tech

- Godot: 4.4.1
- Go: 1.24.3
  - `go version`
  - `go run server/cmd/main.go`
  - debugger: dlv (`go install -v github.com/go-delve/delve/cmd/dlv@latest`)
- Protocol Buffers: 31.1 (network packets)
  - `protoc --version`
- VSCode
  - extensions: `.vscode/extensions.json`
  - settings: `.files/{OS}/.vscode/settings.json` (copy to `.vscode/settings.json`)
  - launch: `.files/{OS}/.vscode/launch.json` (copy to `.vscode/launch.json`)
- Go + protocol buffers
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `protoc -I="shared" --go_out="server" "shared/packets.proto"`
- Godot + protocol buffers
  - https://github.com/oniksan/godobuf
- database: https://sqlc.dev/
  - check
    - convex
    - surrealdb
- websockets
  - TCP based (not for fast paced, realtime game; use UDP otherwise)
  - https://github.com/gorilla/websocket
