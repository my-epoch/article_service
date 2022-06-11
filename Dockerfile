FROM golang:latest

RUN mkdir "build"
ADD . /build
WORKDIR /build

RUN go build cmd/grpc_server/grpc_server.go
EXPOSE 50050
ENTRYPOINT "./grpc_server"