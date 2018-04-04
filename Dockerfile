FROM golang
ADD . /go/src/github.com/wy3148/hello-world
RUN go install github.com/wy3148/hello-world
EXPOSE 3335
ENTRYPOINT ["hello-world"]
