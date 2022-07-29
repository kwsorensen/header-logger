FROM golang:1.17-alpine as build
WORKDIR /app
ADD . .
ENV GOPATH /go

RUN go build .


FROM alpine:latest
COPY --from=build /app/header-logger /app/
WORKDIR /app
RUN chown 65534:65534 header-logger
USER 65534:65534
ENTRYPOINT [ "./header-logger" ]