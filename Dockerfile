FROM golang:1.14.6
WORKDIR /go/src/github.com/transnano/steganography-go/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o steganography -ldflags "-s -w"

FROM alpine:3.12.0
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/transnano/steganography-go/steganography /bin/steganography
ENTRYPOINT ["/bin/steganography"]
