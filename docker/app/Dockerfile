FROM golang:latest
WORKDIR /go/app
ENV GO111MODULE=on
RUN go install github.com/air-verse/air@latest
CMD ["air","-c",".air.toml"]