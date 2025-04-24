FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ../../Users/Carlos%20Carvalho/AppData/Local/Temp/Rar$DRa13832.48183.rartemp .

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN go build -o main ./cmd/server

CMD ["./main"]