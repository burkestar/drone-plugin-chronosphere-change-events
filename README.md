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

To build the distroless image (~4MB):

```bash
make build
```

Usage:

```bash
docker run --rm burkestar/drone-plugin-chronosphere-change-events:latest --help
```

Note that `publish` is the default command used when running the container.

Publish event:

```bash
docker run --rm \
    -e PLUGIN_CHRONOSPHERE_EVENTS_API="${CHRONOSPHERE_EVENTS_API}" \
    -e PLUGIN_CHRONOSPHERE_API_TOKEN="${CHRONOSPHERE_API_TOKEN}" \
    -e PLUGIN_CATEGORY=deploys \
    -e PLUGIN_TYPE=deploy_test \
    -e PLUGIN_SOURCE=local \
    -e PLUGIN_LABELS="environment=local;cluster=none" \
    burkestar/drone-plugin-chronosphere-change-events:latest
```


### As Drone plugin

```yaml
steps:
- name: chronosphere-change-event
  image: burkestar/drone-plugin-chronosphere-change-events
  settings:
    chronosphere_events_api: https://COMPANY.chronosphere.io/api/v1/data/events
    chronosphere_api_token: SOMETHING_SECRET
    category: deploys
    type: deploy_test
    source: drone
    labels: "environment=dev;cluster=dev-cluster"
```

### As plugin in Harness pipeline

Assumptions:

- Harness is configured with a connector to Docker registry (`account.harnessImage`)
- Harness is configured with org-level variable `chronosphere_events_api`
- Harness is configured with org-level secret `chronosphere_api_token`

Within your pipeline YAML:

```yaml
steps:
  - step:
    type: Plugin
    name: chronosphere change event
    identifier: chronosphere_change_event
    spec:
      connectorRef: account.harnessImage ## Docker connector to pull the plugin's Docker image
      image: burkestar/drone-plugin-chronosphere-change-events
      settings:
        chronosphere_events_api: <+variable.org.chronosphere_events_api>
        chronosphere_api_token: <+secrets.getValue("org.chronosphere_api_token")>
        category: deploys
        type: deploy_test
        source: drone
        labels: "environment=dev;cluster=dev-cluster"
```


## Development

Publish image to Docker Hub:

```bash
make publish
```
