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

Usage help:

```bash
drone-plugin-chronosphere-change-events --help

drone-plugin-chronosphere-change-events publish--help
```

Publishing a deploy event:

```bash
drone-plugin-chronosphere-change-events publish --category deploys --type deploy_test --source local --labels 'environment=local;cluster=none'
```

Troubleshooting with `--dry-run` (skips API call) and `--debug` (outputs diagnostics to the console):

```bash
drone-plugin-chronosphere-change-events --dry-run --debug publish --category deploys --type deploy_test --source local --labels 'environment=local;cluster=none'
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
