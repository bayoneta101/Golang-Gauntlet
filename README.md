# Golang Gauntlet

Just a personal proyect to improve in golang by implementing random stuff.


## Architecture

The project is split into two separate microservices — `client` and `server` — each built into its own Docker container with only its own source code included. Both are defined in `compose.yml` and share a single `Dockerfile` that uses a `SERVICE` build arg to select the target. To build both containers at once:

```bash
COMPOSE_BAKE=true docker compose build
```

### Folder structure

```
.
├── proto/          # .proto definitions + go:generate directive
├── gen/            # generated gRPC code (do not edit manually)
├── client/
│   ├── grpc/       # client-side gRPC logic
│   └── main/       # client entrypoint
└── server/
    ├── grpc/       # server-side gRPC logic
    └── main/       # server entrypoint
```

### Generating gRPC code

Proto files live in `proto/` and output goes to `gen/`. To regenerate after changing a `.proto` file:

```bash
go generate ./...
```

## List of implemented features

### gRPC — Greeter service

A simple client/server demo using Protocol Buffers and gRPC. The server exposes a `Hello` RPC that takes a name and returns a greeting. The client calls it once.

To call it manually from the command line (server must be running):

```bash
grpcurl -plaintext -d '{"name": "world"}' localhost:9000 grpc.Greeter/Hello
```

