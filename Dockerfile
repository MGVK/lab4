FROM ubuntu

RUN apt update && apt install -y software-properties-common

RUN add-apt-repository -y ppa:longsleep/golang-backports

RUN apt update && apt install -y golang-go git

RUN mkdir main && cd main

RUN go get -t github.com/lib/pq

COPY src /

RUN go build -v .

EXPOSE 8080