FROM golang:1.13

WORKDIR /go/src/github.com/danimal141/pivotal-go

RUN if [ -f go.mod ]; then cp go.mod .; fi
RUN if [ -f go.sum ]; then cp go.sum .; fi
RUN go mod download

COPY . .
