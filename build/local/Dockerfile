# final runner image
FROM alpine:latest as runner

LABEL maintaner="Mardan M"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
 && apk add tzdata

WORKDIR /opt

COPY ./bin/app-cli /opt/app-cli

EXPOSE 2333

ENTRYPOINT ["/opt/app-cli"]