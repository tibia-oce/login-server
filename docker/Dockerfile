# Stage 1: Download all dependencies
FROM golang:1.23.0-bullseye as builder

WORKDIR /usr/local/go/src/github.com/tibia-oce/login-server
COPY go.mod .
COPY go.sum .
COPY src ./src

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /bin/login-server ./src/

FROM ubuntu:20.04

# Install necessary packages
RUN apt-get update && apt-get install -y \
    iptables \
    # other packages if needed \
    && rm -rf /var/lib/apt/lists/*

# Copy our static executable
COPY --from=builder /bin/login-server /bin/login-server

EXPOSE 80
EXPOSE 8080
EXPOSE 9090

ENTRYPOINT ["/bin/login-server"]