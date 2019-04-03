FROM golang:latest as builder
WORKDIR /go/src/github.com/hdiomede/travel-scanner
RUN go get -u golang.org/x/text/...
RUN go get -u github.com/labstack/echo/...
RUN go get -u  github.com/stretchr/testify/assert/...
RUN go get -u  github.com/stretchr/testify/mock/...
COPY ./ /go/src/github.com/hdiomede/travel-scanner

WORKDIR /go/src/github.com/hdiomede/travel-scanner/testing
RUN go test -v

WORKDIR /go/src/github.com/hdiomede/travel-scanner/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go install

FROM scratch
COPY --from=0 /go/bin/api .
EXPOSE 8080