FROM golang:1.11
MAINTAINER D119 <contact@moonbear.cn>

ARG APP_HOME=/app

COPY . $APP_HOME

ENV GOPROXY=https://goproxy.io \
    PORT=1234

RUN cd $APP_HOME \
    && go build -o /root/api-server .
    && rm -rf $GOPATH/pkg/ $APP_HOME /root/.cache

EXPOSE $PORT

ENTRYPOINT /root/api-server -p $PORT