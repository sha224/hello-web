FROM golang:1.14
WORKDIR /hello-web
COPY main.go .
RUN go build --ldflags '-s -w -linkmode "external" -extldflags "-static"' -o hello-web .

FROM scratch
COPY --from=0 /hello-web/hello-web .
CMD ["./hello-web"]