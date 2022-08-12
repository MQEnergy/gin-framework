FROM golang:latest
LABEL maintainer="MQEnergy Developers <bbxycx18@gmail.com>" version="1.0" license="MIT"

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/chenxi2015/gin-framework
COPY . $GOPATH/src/github.com/chenxi2015/gin-framework
RUN go build .

EXPOSE 9828
ENTRYPOINT ["./gin-framework"]