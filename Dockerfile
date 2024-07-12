FROM bitnami/golang:1.22.5

WORKDIR /go/src/echo
ENV GOPROXY=https://proxy.golang.org

COPY . .
RUN go mod download && go mod verify
RUN go build -v -o . ./...

CMD ["./echo"]
