FROM daocloud.io/library/golang:1.9

ENV APP_DIR $GOPATH/src/kjlive-service
RUN mkdir -p $APP_DIR/logs

# Set the entrypoint
ENTRYPOINT (cd $GOPATH/src/kjlive-service && ./kjlive-service)
ADD . $APP_DIR
