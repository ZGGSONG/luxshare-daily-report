FROM alpine:latest

RUN mkdir "/app"
WORKDIR "/app"

RUN mkdir "/app/config"
RUN mkdir "/app/log"
RUN mkdir "/app/upload"

# 修正时区
RUN apk update && apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone


COPY ksat-mrsb "/app/ksat-mrsb"

EXPOSE 7201

ENTRYPOINT ["./ksat-mrsb"]