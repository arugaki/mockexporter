FROM daocloud.io/golang
RUN mkdir /gopath
ENV GOPATH /gopath
ADD . /gopath/
RUN cd /gopath/src/main/ && go build main.go && mv main /go && rm -rf /gopath
EXPOSE 9120
CMD ["./main"]

