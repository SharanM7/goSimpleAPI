From golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/SharanM7/goSimpleAPI
RUN cd /build && git clone https://github.com/SharanM7/goSimpleAPI.git

RUN cd /build/main && go build

EXPOSE 8090

ENTRYPOINT ["/build/main"]