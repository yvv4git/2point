FROM golang:latest 

#ADD . /go/src/github.com/yvv4git/2endpoint
#WORKDIR /go/src/github.com/yvv4git/2endpoint

#RUN go get github.com/sarulabs/di \
#&& mkdir -p /go/src/github.com/yvv4git/2endpoint \
#&& go build -o build/server.bin app/cmd/main.go

#CMD ["build/server.bin"]

RUN mkdir -p /app
ADD . /app
WORKDIR /app
RUN go build -mod=vendor -o build/server.bin app/cmd/main.go

CMD ["build/server.bin"]

