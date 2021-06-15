################################################################################
## GO BUILDER
################################################################################
FROM golang:1.16.5 as builder

ENV VERSION 1.0.0
ENV BUILD_DIR /build

WORKDIR ${BUILD_DIR}

COPY go.* ./
RUN go mod download
COPY . .

RUN go test -cover ./...
RUN CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo -ldflags "-X github.com/halverneus/https-redirec/lib/cli/version.version=${VERSION}" -o /redirect /build/bin/redirect

RUN adduser --system --no-create-home --uid 1000 --shell /usr/sbin/nologin static

################################################################################
## DEPLOYMENT CONTAINER
################################################################################
FROM scratch

EXPOSE 8080
COPY --from=builder /redirect /
COPY --from=builder /etc/passwd /etc/passwd

USER static
ENTRYPOINT ["/redirect"]
CMD []

# Metadata
LABEL life.apets.vendor="Halverneus" \
    life.apets.url="https://github.com/halverneus/https-redirect" \
    life.apets.name="HTTPS Redirect Server" \
    life.apets.description="A tiny HTTPS redirect server" \
    life.apets.version="v1.0.0" \
    life.apets.schema-version="1.0"
