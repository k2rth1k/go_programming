FROM golang:latest

# System setup
RUN apt-get update
RUN apt install -y protobuf-compiler
RUN protoc --version

RUN git clone https://github.com/k2rth1k/go_programming.git
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
RUN go get -u github.com/golang/protobuf/protoc-gen-go
WORKDIR go_programming
RUN pwd; ls -l;
EXPOSE 50051

RUN make go
