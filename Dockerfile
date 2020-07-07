FROM golang:alpine
WORKDIR /hello-web
COPY main.go .
RUN CGO_ENABLED=0 go build -ldflags "-s -w"

FROM scratch
COPY --from=0 /hello-web/hello-web .
CMD ["./hello-web"]