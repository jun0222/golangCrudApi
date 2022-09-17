# このあたりを修正したときはdocker-compose build必要
FROM golang:alpine3.15

WORKDIR /api

COPY ./apiSample/endpoint.go .

RUN go mod init todo && go mod tidy

# .に./apiSample/endpoint.goをCOPYした後なのでendpoint.goと相対パスを指定
CMD ["go", "run", "endpoint.go"]