From golang:latest

RUN mkdir /build
WORKDIR /build

RUN cd /build 

RUN ./main

EXPOSE 8080

ENTRYPOINT ["/build/main"]