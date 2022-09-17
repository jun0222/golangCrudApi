# このあたりを修正したときはdocker-compose build必要
FROM golang:alpine3.15

WORKDIR /api

# helloWorld用
# COPY ./apiSample/endpoint.go .

# todoアプリのcrud用（gin使用）
COPY ./crud/main.go .

RUN go mod init todo && go mod tidy

# .に./apiSample/endpoint.goをCOPYした後なのでendpoint.goと相対パスを指定
# helloWorld用
# CMD ["go", "run", "endpoint.go"] 

# todoアプリのcrud用（gin使用）
CMD ["go", "run", "main.go"]