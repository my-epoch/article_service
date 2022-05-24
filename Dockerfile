FROM golang:latest

RUN mkdir "build"
ADD . /build
WORKDIR /build

RUN go build -o "build/grpc_server" cmd/grpc_server/grpc_server.go
EXPOSE 50050
ENTRYPOINT ["build/grpc_server"]