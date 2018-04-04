# hello-world example

## prerequisite
   docker
   golang
## code
  go get github.com/wy3148/hello-world
  cd $GOPATH/src/github.com/wy3148/hello-world
  go test -v
  docker build . -t hello-world
  docker run --rm -p 3000:3333 hello-world
  
