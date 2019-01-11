FROM golang:1.11.2-alpine3.8
WORKDIR /Go-docker
COPY dist/Go-docker /bin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
CMD ["./app"]