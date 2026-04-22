# syntax=docker/dockerfile:1
ARG DEPLOIO_GIT_REVISION="0.0.1-dockerbuild"
FROM golang:1.26 AS build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go test -v

ARG DEPLOIO_GIT_REVISION
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags="-X 'main.appVersion=${DEPLOIO_GIT_REVISION}'"

FROM gcr.io/distroless/static-debian12

COPY --from=build /go/bin/app /
CMD ["/app"]
