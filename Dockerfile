# syntax=docker/dockerfile:1
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/aleskeinovikov/playground/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/aleskeinovikov/playground/app ./
EXPOSE 8080
CMD ["./app"]  
