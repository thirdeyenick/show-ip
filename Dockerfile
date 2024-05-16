# syntax=docker/dockerfile:1
FROM golang:1.21 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags="-X 'main.appVersion=0.0.1-dockerbuild'"

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]
