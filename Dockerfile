FROM golang:1.18.1
WORKDIR /temp

RUN apt-get update && apt-get install -y unzip
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-linux-x86_64.zip -O protobuf.zip && \
    unzip -o protobuf.zip -d protobuf && \
    chmod -R 755 protobuf/*

ENV PATH $PATH:/temp/protobuf/bin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2