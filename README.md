# Golang Gauntlet

Just a personal proyect to improve in golang by implementing random stuff.


## Architecture

The project is split into two separate microservices — `client` and `server` — each built into its own Docker container with only its own source code included. Both are defined in `compose.yml` and share a single `Dockerfile` that uses a `SERVICE` build arg to select the target. To build both containers at once:

```bash
COMPOSE_BAKE=true docker compose build
```

## List of implemented features

Nothing for now :).

