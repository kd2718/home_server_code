FROM golang:1.24-alpine AS builder

WORKDIR /project/go-server/

COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o /project/go-server/build/app .

FROM alpine:latest

COPY --from=builder /project/go-server/build/app /project/go-server/build/app

EXPOSE 8081
ENTRYPOINT [ "/project/go-server/build/app" ]
