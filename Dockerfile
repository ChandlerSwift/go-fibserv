FROM golang:1.14 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build ./fibserv.go

FROM scratch

COPY --from=build /build/fibserv /entrypoint

CMD ["/entrypoint"]
