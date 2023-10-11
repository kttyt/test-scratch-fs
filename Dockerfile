ARG TARGETOS=linux
ARG TARGETARCH=arm64

FROM golang:1.20-rc-alpine as builder
ENV GOCACHE=/go_cache \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH}
RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY go.mod ./
RUN go mod download

#No cache after this action
COPY . .
RUN go build -ldflags="-w -s" -o app .


FROM scratch
COPY --from=builder /src/app /app
COPY index.html /index.html
ENTRYPOINT ["/app"]
