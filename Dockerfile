FROM golang:1.22.3-bullseye as base
# Distroless based on https://klotzandrew.com/blog/smallest-golang-docker-image
# Debian Bullseye is used as the build container, with the runtime container based on scratch
# with the compiled binary copied inside.

# Create a non-root user
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid 65532 \
    drone-plugin-user

WORKDIR $GOPATH/src/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /main .

USER drone-plugin-user:drone-plugin-user

ENTRYPOINT ["/main", "publish"]
