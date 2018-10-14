FROM golang:1.11
MAINTAINER D119 <contact@moonbear.cn>

ARG APP_HOME=/app
COPY . $APP_HOME

RUN export GOPATH=$GOPATH:$APP_HOME \
    && cd $APP_HOME \
    && go get github.com/themoonbear/gvt \
    && gvt restore -g \
    && go build -o /root/proxy ./src \
    && rm -rf $GOPATH $APP_HOME /root/.cache

EXPOSE 1324

ENTRYPOINT /root/proxy -p 1324