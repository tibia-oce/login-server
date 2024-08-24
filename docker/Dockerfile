FROM golang:1.16.4-alpine as builder
WORKDIR /login-server
COPY go.mod .
COPY go.sum .
COPY src ./src
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /bin/login-server ./src/


FROM alpine:latest

COPY --from=builder /bin/login-server /bin/login-server
ENTRYPOINT ["/bin/login-server"]
