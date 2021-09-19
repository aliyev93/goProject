FROM golang:1.17 as builder
WORKDIR /build
ADD * ./

RUN make build
RUN pwd && ls -la

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/app .
COPY config.yaml .
RUN adduser gouser --disabled-password
USER gouser
CMD ["./app"]