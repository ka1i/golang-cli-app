FROM golang:1.20 AS builder

ENV GOPROXY https://proxy.golang.com.cn,direct

WORKDIR /golang-cli-app/

COPY . .

RUN go mod tidy
RUN make build

# final runner image
FROM alpine:latest as runner

LABEL maintaner="Mardan M"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
 && apk add tzdata

WORKDIR /opt

COPY --from=builder /golang-cli-app/bin/app-cli /opt/app-cli

EXPOSE 2333

ENTRYPOINT ["/opt/app-cli"]