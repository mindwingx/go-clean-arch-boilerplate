FROM golang:1.19

WORKDIR /app

COPY src .

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

# RUN go test -v .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main serve/main.go
RUN chmod +x main

CMD ["/app/main"]