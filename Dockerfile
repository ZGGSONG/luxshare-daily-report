FROM alpine:latest

RUN mkdir "/app"
WORKDIR "/app"

RUN mkdir "/app/config"
RUN mkdir "/app/log"
RUN mkdir "/app/upload"

COPY receive-files "/app/receive-files"

EXPOSE 7201

ENTRYPOINT ["./receive-files"]