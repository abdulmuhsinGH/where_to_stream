FROM golang:1.13.2-alpine3.10

RUN mkdir -p go/src/ecormmerce_api

WORKDIR /go/src/ecormmerce_api

COPY . /go/src/ecormmerce_api

RUN go install ecormmerce_api

CMD /go/bin/ecormmerce_api

EXPOSE 8080