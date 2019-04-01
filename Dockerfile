FROM golang:latest as builder
WORKDIR /go/src/github.com/hdiomede/travel-scanner
RUN go get -u golang.org/x/text/...
RUN go get -u github.com/labstack/echo/...
COPY ./ /go/src/github.com/hdiomede/travel-scanner

WORKDIR /go/src/github.com/hdiomede/travel-scanner/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go install

FROM scratch
COPY --from=0 /go/bin/api .
COPY ./cmd/api/file.csv /
EXPOSE 8080
CMD ["./api"]