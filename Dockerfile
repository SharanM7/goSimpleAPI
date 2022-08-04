FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

# RUN export GO111MODULE=on
# RUN go get github.com/SharanM7/goSimpleAPI
# RUN git clone https://github.com/SharanM7/goSimpleAPI.git

RUN go build main.go 

# EXPOSE 8090

CMD ["/app/main"]