FROM golang:1.24 AS builder

WORKDIR /jester

COPY go.mod go.sum api/go.mod api/go.sum ./

RUN go mod download

COPY . .

RUN make build

FROM debian:bookworm-slim

# bookworm-slim does not have HOME set by default
ENV HOME=/home/jester
RUN mkdir -p $HOME && chown -R 2000:2000 $HOME

USER 2000:2000

COPY --from=builder /jester/build/jesterd /usr/local/bin/jesterd

# Assumes default ports:
#   9091 for GRPC
#   2112 for METRICS
EXPOSE 9091 2112

USER 2000:2000

ENTRYPOINT ["/usr/local/bin/jesterd"]
