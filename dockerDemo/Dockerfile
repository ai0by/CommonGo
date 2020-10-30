FROM golang

ENV TZ Asia/Shanghai

ENV GO111MODULE on

ENV GOPROXY https://goproxy.io

MAINTAINER ai0by


#RUN go get github.com/beego/bee

ENV GOPROXY https://goproxy.io/

ENV GO111MODULE on


ADD . .

#RUN go mod vendor

EXPOSE 8100

#ENV GO111MODULE off

#CMD bee run

CMD go run main.go