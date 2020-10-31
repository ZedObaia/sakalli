FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o sakalli-server .

WORKDIR /dist

RUN cp /build/sakalli-server .

FROM scratch

COPY --from=builder /dist/sakalli-server /

EXPOSE 8080

ENTRYPOINT ["/sakalli-server"]