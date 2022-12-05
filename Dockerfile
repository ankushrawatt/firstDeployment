FROM golang:1.16-alpine
RUN apk update
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
RUN go build -o /main
CMD [ "/main.go" ]