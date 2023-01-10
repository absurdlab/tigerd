FROM golang:1.18.1 AS buildStage

ENV HOME /build
ENV CGO_ENABLED 0
ENV GOOS linux

ARG BUILD_VERSION

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build \
    -a \
    -ldflags "-X github.com/absurdlab/tigerd/buildinfo.Version=$BUILD_VERSION" \
    -installsuffix cgo \
    -o tigerd \
    .

FROM alpine:3.15

LABEL org.opencontainers.image.title="tigerd"
LABEL org.opencontainers.image.source="https://github.com/absurdlab/tigerd"
LABEL org.opencontainers.image.authors="Weinan Qiu"

COPY --from=buildStage /build/tigerd /usr/bin/tigerd

ENTRYPOINT ["/usr/bin/tigerd"]