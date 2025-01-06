# syntax=docker/dockerfile:1
ARG APP_VERSION="0.0.1-dockerbuild"
FROM golang:1.21 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go test -v

ARG APP_VERSION
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags="-X 'main.appVersion=${APP_VERSION}'"

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]
