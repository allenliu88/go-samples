FROM --platform=$TARGETPLATFORM golang:1.20.10-alpine3.18 as go-builder

RUN apk add --no-cache gcc g++

## 目录
RUN mkdir -p /opt/repo/go-samples
WORKDIR /opt/repo/go-samples

## 拷贝文件
COPY . ./

## 设置代理
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

## 交叉编译
RUN go mod verify && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/go-samples-linux-amd64 ./main.go

# Final stage
FROM --platform=$TARGETPLATFORM alpine:3.18

LABEL maintainer="Grafana team <hello@grafana.com>"

ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

WORKDIR /opt/repo/go-samples

COPY --from=go-builder /opt/repo/go-samples/bin/go-samples-linux-amd64 ./bin/

EXPOSE 8080
ENTRYPOINT [ "/opt/repo/go-samples/bin/go-samples-linux-amd64" ]