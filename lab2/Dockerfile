FROM golang:1.18 as builder

RUN apt-get update
# RUN apt-get install -y protobuf-compiler
# RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /project

COPY Makefile /project

# COPY go.mod /project
# RUN go mod download

COPY . /project

RUN make bin/server

#RUN apt-get update && apt-get install -y python3
# RUN apt-get install dialog apt-utils -y

# RUN apt-get update && apt-get install -y iputils-ping


FROM ubuntu:20.04

RUN apt update && yes | apt upgrade
RUN yes | apt install python3 iputils-ping

WORKDIR /

COPY --from=builder /project/bin/server /bin/server

ENTRYPOINT ["/bin/server"]
