FROM alpine:latest

RUN mkdir "/app"
WORKDIR "/app"

RUN mkdir "/app/config"
RUN mkdir "/app/log"
RUN mkdir "/app/upload"

COPY ksat-mrsb "/app/ksat-mrsb"

EXPOSE 7201

ENTRYPOINT ["./ksat-mrsb"]