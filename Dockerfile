FROM alpine:latest

COPY main /opt/main
#RUN chmod +x /opt/main

ENTRYPOINT /opt/main