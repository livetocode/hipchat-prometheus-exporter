FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY ./build/app hipchat-prometheus-exporter

ENTRYPOINT ["./hipchat-prometheus-exporter"]  
