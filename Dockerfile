FROM golang:1.18.1 AS buildStage

ARG BUILD_VERSION
ARG BUILD_TIME

ENV HOME /build
ENV CGO_ENABLED 0
ENV GOOS linux
ENV BUILD_VERSION=$BUILD_VERSION
ENV BUILD_TIME=$BUILD_TIME

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build \
    -a \
    -ldflags "-w -s -X absurdlab.io/tigerd/internal/buildinfo.Version=$BUILD_VERSION -X absurdlab.io/tigerd/internal/buildinfo.CompiledAt=$BUILD_TIME" \
    -installsuffix cgo \
    -o tigerd \
    .

FROM alpine:3.15

COPY --from=buildStage /build/tigerd /usr/bin/tigerd

ENTRYPOINT ["/usr/bin/tigerd"]