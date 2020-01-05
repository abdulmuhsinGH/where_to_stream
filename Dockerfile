FROM golang:1.13.2-alpine3.10

RUN mkdir -p go/src/where_to_stream

WORKDIR /go/src/where_to_stream

COPY . /go/src/where_to_stream

RUN go install where_to_stream

CMD /go/bin/where_to_stream

EXPOSE 8080