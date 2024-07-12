FROM bitnami/golang:1.22.5 AS builder

WORKDIR /go/src/echo
ENV GOPROXY=https://proxy.golang.org

COPY go.mod go.sum main.go ./
RUN go mod download && go mod verify

COPY config ./
COPY controllers ./
COPY middleware ./
COPY models ./
COPY routes ./
COPY views ./
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
