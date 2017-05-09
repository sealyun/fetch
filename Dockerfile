FROM dev.reg.iflytek.com/base/golang:1.8.0
RUN mkdir -p $GOPATH/src/fetch/ && go get github.com/tools/godep
COPY . $GOPATH/src/fetch/
RUN set -ex && cd $GOPATH/src/fetch/ && godep go build -o main && pwd && ls
CMD $GOPATH/src/fetch/main moby/moby
