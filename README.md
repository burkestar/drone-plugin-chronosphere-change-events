# Drone plugin for Chronosphere Change Events

[Drone plugin](https://docs.drone.io/plugins/overview/) for Harness pipelines to publish [Chronosphere Change Events](https://docs.chronosphere.io/investigate/change-events).

## Usage

### On host

Setup:

- golang installed (`1.22.3`)
- `GOBIN` environment variable configured where go binaries will be installed
- `GOBIN` on your `PATH` so installed go binaries are discovered

Install:

```bash
go install
```

Use:

```bash
drone-plugin-chronosphere-change-events
```

Usage help:

```bash
drone-plugin-chronosphere-change-events --help
```


### In Docker container

Build container image:

```bash
make build
```

Use:

```bash
make run
```


## Development

Publish image to Docker Hub:

```bash
make publish
```
