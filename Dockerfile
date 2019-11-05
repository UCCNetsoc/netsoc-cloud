FROM golang:1.13

ENV GO111MODULE=on

WORKDIR /netsoc-cloud

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

EXPOSE 7070

ENTRYPOINT ./cloud