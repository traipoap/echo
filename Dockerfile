FROM bitnami/golang:1.22 AS builder

WORKDIR /go/src/projec
ENV GOPROXY=https://proxy.golang.org

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY server.go .
RUN go build -v -o /usr/local/bin/app ./...

FROM amazoncorretto:21.0.3-alpine3.19

WORKDIR /echo

COPY --from=builder /usr/local/bin/app /echo

CMD ["app"]
